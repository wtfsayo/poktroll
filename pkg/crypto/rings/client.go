package rings

import (
	"context"
	"fmt"

	"cosmossdk.io/depinject"
	ring_secp256k1 "github.com/athanorlabs/go-dleq/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/noot/ring-go"

	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/crypto"
	pubkeyclient "github.com/pokt-network/poktroll/pkg/crypto/pubkey_client"
	"github.com/pokt-network/poktroll/pkg/polylog"
	"github.com/pokt-network/poktroll/x/service/types"
)

var _ crypto.RingClient = (*ringClient)(nil)

type ringClient struct {
	// logger is the logger for the ring cache.
	logger polylog.Logger

	// applicationQuerier is the querier for the application module, and is
	// used to get the addresses of the gateways an application is delegated to.
	applicationQuerier client.ApplicationQueryClient

	// pubKeyClient
	pubKeyClient crypto.PubKeyClient
}

// NewRingClient returns a new ring client constructed from the given dependencies.
// It returns an error if the required dependencies are not supplied.
//
// Required dependencies:
// - polylog.Logger
// - client.ApplicationQueryClient
// - client.AccountQueryClient
func NewRingClient(deps depinject.Config) (_ crypto.RingClient, err error) {
	rc := new(ringClient)
	rc.pubKeyClient, err = pubkeyclient.NewPubKeyClient(deps)
	if err != nil {
		return nil, err
	}

	if err := depinject.Inject(
		deps,
		&rc.logger,
		&rc.applicationQuerier,
	); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetRingForAddress returns the ring for the address provided. The ring is created by
// querying for the application address and delegated gateways' account public keys and
// converting them to their secp256k1 curve points.
func (rc *ringClient) GetRingForAddress(
	ctx context.Context,
	appAddress string,
) (*ring.Ring, error) {
	pubKeys, err := rc.getDelegatedPubKeysForAddress(ctx, appAddress)
	if err != nil {
		return nil, err
	}
	// Cache the ring's points for future use
	rc.logger.Debug().
		Str("app_address", appAddress).
		Msg("updating ring ringsByAddr for app")

	// Get the points on the secp256k1 curve for the public keys in the ring.
	points, err := pointsFromPublicKeys(pubKeys...)
	if err != nil {
		return nil, err
	}

	// Return the ring the constructed from the public key points on the secp256k1 curve.
	return newRingFromPoints(points)
}

// VerifyRelayRequestSignature verifies the relay request signature against the
// ring for the application address in the relay request.
func (rc *ringClient) VerifyRelayRequestSignature(
	ctx context.Context,
	relayRequest *types.RelayRequest,
) error {
	rc.logger.Debug().
		Fields(map[string]any{
			"session_id":          relayRequest.Meta.SessionHeader.SessionId,
			"application_address": relayRequest.Meta.SessionHeader.ApplicationAddress,
			"service_id":          relayRequest.Meta.SessionHeader.Service.Id,
		}).
		Msg("verifying relay request signature")

	// Extract the relay request's ring signature
	if relayRequest.Meta == nil {
		return ErrRingClientEmptyRelayRequestSignature.Wrapf(
			"request payload: %s", relayRequest.Payload,
		)
	}

	signature := relayRequest.Meta.Signature
	if signature == nil {
		return ErrRingClientInvalidRelayRequest.Wrapf(
			"missing signature from relay request: %v", relayRequest,
		)
	}

	ringSig := new(ring.RingSig)
	if err := ringSig.Deserialize(ring_secp256k1.NewCurve(), signature); err != nil {
		return ErrRingClientInvalidRelayRequestSignature.Wrapf(
			"error deserializing ring signature: %v", err,
		)
	}

	if relayRequest.Meta.SessionHeader.ApplicationAddress == "" {
		return ErrRingClientInvalidRelayRequest.Wrap(
			"missing application address from relay request",
		)
	}

	// Get the ring for the application address of the relay request
	appAddress := relayRequest.Meta.SessionHeader.ApplicationAddress
	appRing, err := rc.GetRingForAddress(ctx, appAddress)
	if err != nil {
		return ErrRingClientInvalidRelayRequest.Wrapf(
			"error getting ring for application address %s: %v", appAddress, err,
		)
	}

	// Verify the ring signature against the ring
	if !ringSig.Ring().Equals(appRing) {
		return ErrRingClientInvalidRelayRequestSignature.Wrapf(
			"ring signature does not match ring for application address %s", appAddress,
		)
	}

	// Get and hash the signable bytes of the relay request
	requestSignableBz, err := relayRequest.GetSignableBytesHash()
	if err != nil {
		return ErrRingClientInvalidRelayRequest.Wrapf("error getting signable bytes: %v", err)
	}

	// Verify the relay request's signature
	if valid := ringSig.Verify(requestSignableBz); !valid {
		return ErrRingClientInvalidRelayRequestSignature.Wrapf("invalid ring signature")
	}
	return nil
}

// getDelegatedPubKeysForAddress returns the ring used to sign a message for
// the given application address, by querying the application module for it's
// delegated pubkeys and converting them to points on the secp256k1 curve in
// order to create the ring.
func (rc *ringClient) getDelegatedPubKeysForAddress(
	ctx context.Context,
	appAddress string,
) ([]cryptotypes.PubKey, error) {
	// Get the application's on chain state.
	app, err := rc.applicationQuerier.GetApplication(ctx, appAddress)
	if err != nil {
		return nil, err
	}

	// Create a slice of addresses for the ring.
	ringAddresses := make([]string, 0)
	ringAddresses = append(ringAddresses, appAddress) // app address is index 0
	if len(app.DelegateeGatewayAddresses) == 0 {
		// add app address twice to make the ring size of mininmum 2
		// TODO_HACK: We are adding the appAddress twice because a ring
		// signature requires AT LEAST two pubKeys. When the Application has
		// not delegated to any gateways, we add the application's own address
		// twice. This is a HACK and should be investigated as to what is the
		// best approach to take in this situation.
		ringAddresses = append(ringAddresses, appAddress)
	} else {
		// add the delegatee gateway addresses
		ringAddresses = append(ringAddresses, app.DelegateeGatewayAddresses...)
	}

	rc.logger.Debug().
		// TODO_TECHDEBT: implement and use `polylog.Event#Strs([]string)` instead of formatting here.
		Str("addresses", fmt.Sprintf("%v", ringAddresses)).
		Msg("converting addresses to points")

	return rc.addressesToPubKeys(ctx, ringAddresses)
}

// addressesToPubKeys uses the public key client to query the account module for
// the public key corresponding to each address given.
func (rc *ringClient) addressesToPubKeys(
	ctx context.Context,
	addresses []string,
) ([]cryptotypes.PubKey, error) {
	pubKeys := make([]cryptotypes.PubKey, len(addresses))
	for i, addr := range addresses {
		acc, err := rc.pubKeyClient.GetPubKeyFromAddress(ctx, addr)
		if err != nil {
			return nil, err
		}
		pubKeys[i] = acc
	}
	return pubKeys, nil
}