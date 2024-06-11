package keeper_test

import (
	"testing"

	"cosmossdk.io/depinject"
	ring_secp256k1 "github.com/athanorlabs/go-dleq/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pokt-network/ring-go"
	"github.com/pokt-network/smt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pokt-network/poktroll/pkg/crypto/rings"
	"github.com/pokt-network/poktroll/pkg/polylog/polyzero"
	"github.com/pokt-network/poktroll/pkg/relayer"
	"github.com/pokt-network/poktroll/pkg/relayer/protocol"
	keepertest "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/testutil/keeper/common"
	"github.com/pokt-network/poktroll/testutil/testkeyring"
	"github.com/pokt-network/poktroll/x/proof/keeper"
	"github.com/pokt-network/poktroll/x/proof/types"
	servicetypes "github.com/pokt-network/poktroll/x/service/types"
	"github.com/pokt-network/poktroll/x/shared"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

// TODO_TECHDEBT(@bryanchriswhite): Simplify this file; https://github.com/pokt-network/poktroll/pull/417#pullrequestreview-1958582600

const (
	supplierUid = "supplier"
)

var (
	blockHeaderHash         []byte
	expectedMerkleProofPath []byte

	// testProofParams sets:
	//  - the minimum relay difficulty bits to zero so that these tests don't need to mine for valid relays.
	//  - the proof request probability to 1 so that all test sessions require a proof.
	testProofParams = types.Params{
		MinRelayDifficultyBits:  0,
		ProofRequestProbability: 1,
	}
)

func init() {
	// The CometBFT header hash is 32 bytes: https://docs.cometbft.com/main/spec/core/data_structures
	blockHeaderHash = make([]byte, 32)
	expectedMerkleProofPath = keeper.GetPathForProof(blockHeaderHash, "TODO_BLOCKER_session_id_currently_unused")
}

func TestMsgServer_SubmitProof_Success(t *testing.T) {
	tests := []struct {
		desc              string
		getProofMsgHeight func(
			sharedParams *sharedtypes.Params,
			queryHeight int64,
		) int64
	}{
		{
			desc:              "proof message height equals proof window open height",
			getProofMsgHeight: shared.GetProofWindowOpenHeight,
		},
		{
			desc:              "proof message height equals proof window close height",
			getProofMsgHeight: shared.GetProofWindowCloseHeight,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			opts := []keepertest.ProofKeepersOpt{
				// Set block hash so we can have a deterministic expected on-chain proof requested by the protocol.
				common.WithBlockHash[
					keepertest.ProofKeepersOpt,
					*keepertest.ProofModuleKeepers,
				](blockHeaderHash),
				// Set block height to 1 so there is a valid session on-chain.
				common.WithBlockHeight[
					keepertest.ProofKeepersOpt,
					*keepertest.ProofModuleKeepers,
				](1),
			}
			keepers, ctx := keepertest.NewProofModuleKeepers(t, opts...)
			sharedParams := keepers.SharedKeeper.GetParams(ctx)
			sdkCtx := cosmostypes.UnwrapSDKContext(ctx)

			// Set proof keeper params to disable relaymining and always require a proof.
			err := keepers.Keeper.SetParams(ctx, testProofParams)
			require.NoError(t, err)

			// Construct a keyring to hold the keypairs for the accounts used in the test.
			keyRing := keyring.NewInMemory(keepers.Codec)

			// Create a pre-generated account iterator to create accounts for the test.
			preGeneratedAccts := testkeyring.PreGeneratedAccounts()

			// Create accounts in the account keeper with corresponding keys in the
			// keyring for the application and supplier.
			supplierAddr := testkeyring.CreateOnChainAccount(
				ctx, t,
				supplierUid,
				keyRing,
				keepers,
				preGeneratedAccts,
			).String()
			appAddr := testkeyring.CreateOnChainAccount(
				ctx, t,
				"app",
				keyRing,
				keepers,
				preGeneratedAccts,
			).String()

			service := &sharedtypes.Service{Id: testServiceId}

			// Add a supplier and application pair that are expected to be in the session.
			keepers.AddServiceActors(ctx, t, service, supplierAddr, appAddr)

			// Get the session for the application/supplier pair which is expected
			// to be claimed and for which a valid proof would be accepted.
			// Given the setup above, it is guaranteed that the supplier created
			// will be part of the session.
			sessionHeader := common.GetSessionHeader(ctx, t, keepers, appAddr, service, 1)

			// Construct a proof message server from the proof keeper.
			proofMsgServer := keeper.NewMsgServerImpl(*keepers.Keeper)

			// Prepare a ring client to sign & validate relays.
			ringClient, err := rings.NewRingClient(depinject.Supply(
				polyzero.NewLogger(),
				types.NewAppKeeperQueryClient(keepers.ApplicationKeeper),
				types.NewAccountKeeperQueryClient(keepers.AccountKeeper),
				types.NewSharedKeeperQueryClient(keepers.SharedKeeper),
			))
			require.NoError(t, err)

			// Submit the corresponding proof.
			numRelays := uint(5)
			sessionTree := common.NewFilledSessionTree(
				ctx, t,
				numRelays,
				supplierUid, supplierAddr,
				sessionHeader, sessionHeader, sessionHeader,
				keyRing,
				ringClient,
			)

			// Advance the block height to the test claim msg height.
			claimMsgHeight := shared.GetClaimWindowOpenHeight(
				&sharedParams,
				sessionHeader.GetSessionEndBlockHeight(),
			)
			sdkCtx = sdkCtx.WithBlockHeight(claimMsgHeight)
			ctx = sdkCtx

			// Create a valid claim.
			keepers.CreateClaimAndStoreBlockHash(
				ctx, t, 1,
				supplierAddr,
				appAddr,
				service,
				sessionTree,
				sessionHeader,
			)

			// Advance the block height to the test proof msg height.
			proofMsgHeight := test.getProofMsgHeight(&sharedParams, sessionHeader.GetSessionEndBlockHeight())
			sdkCtx = sdkCtx.WithBlockHeight(proofMsgHeight)
			ctx = sdkCtx

			proofMsg := common.NewTestProofMsg(t,
				supplierAddr,
				sessionHeader,
				sessionTree,
				expectedMerkleProofPath,
			)
			submitProofRes, err := proofMsgServer.SubmitProof(ctx, proofMsg)
			require.NoError(t, err)
			require.NotNil(t, submitProofRes)

			proofRes, err := keepers.AllProofs(ctx, &types.QueryAllProofsRequest{})
			require.NoError(t, err)

			proofs := proofRes.GetProofs()
			require.Lenf(t, proofs, 1, "expected 1 proof, got %d", len(proofs))
			require.Equal(t, proofMsg.SessionHeader.SessionId, proofs[0].GetSessionHeader().GetSessionId())
			require.Equal(t, proofMsg.SupplierAddress, proofs[0].GetSupplierAddress())
			require.Equal(t, proofMsg.SessionHeader.GetSessionEndBlockHeight(), proofs[0].GetSessionHeader().GetSessionEndBlockHeight())
		})
	}
}

func TestMsgServer_SubmitProof_Error_OutsideOfWindow(t *testing.T) {
	opts := []keepertest.ProofKeepersOpt{
		// Set block hash so we can have a deterministic expected on-chain proof requested by the protocol.
		common.WithBlockHash[
			keepertest.ProofKeepersOpt,
			*keepertest.ProofModuleKeepers,
		](blockHeaderHash),
		// Set block height to 1 so there is a valid session on-chain.
		common.WithBlockHeight[
			keepertest.ProofKeepersOpt,
			*keepertest.ProofModuleKeepers,
		](1),
	}
	keepers, ctx := keepertest.NewProofModuleKeepers(t, opts...)

	// Set proof keeper params to disable relaymining and always require a proof.
	err := keepers.Keeper.SetParams(ctx, testProofParams)
	require.NoError(t, err)

	// Construct a keyring to hold the keypairs for the accounts used in the test.
	keyRing := keyring.NewInMemory(keepers.Codec)

	// Create a pre-generated account iterator to create accounts for the test.
	preGeneratedAccts := testkeyring.PreGeneratedAccounts()

	// Create accounts in the account keeper with corresponding keys in the keyring for the application and supplier.
	supplierAddr := testkeyring.CreateOnChainAccount(
		ctx, t,
		supplierUid,
		keyRing,
		keepers,
		preGeneratedAccts,
	).String()
	appAddr := testkeyring.CreateOnChainAccount(
		ctx, t,
		"app",
		keyRing,
		keepers,
		preGeneratedAccts,
	).String()

	service := &sharedtypes.Service{Id: testServiceId}

	// Add a supplier and application pair that are expected to be in the session.
	keepers.AddServiceActors(ctx, t, service, supplierAddr, appAddr)

	// Get the session for the application/supplier pair which is expected
	// to be claimed and for which a valid proof would be accepted.
	// Given the setup above, it is guaranteed that the supplier created
	// will be part of the session.
	sessionHeader := common.GetSessionHeader(ctx, t, keepers, appAddr, service, 1)

	// Construct a proof message server from the proof keeper.
	proofMsgServer := keeper.NewMsgServerImpl(*keepers.Keeper)

	// Prepare a ring client to sign & validate relays.
	ringClient, err := rings.NewRingClient(depinject.Supply(
		polyzero.NewLogger(),
		types.NewAppKeeperQueryClient(keepers.ApplicationKeeper),
		types.NewAccountKeeperQueryClient(keepers.AccountKeeper),
		types.NewSharedKeeperQueryClient(keepers.SharedKeeper),
	))
	require.NoError(t, err)

	// Submit the corresponding proof.
	numRelays := uint(5)
	sessionTree := common.NewFilledSessionTree(
		ctx, t,
		numRelays,
		supplierUid, supplierAddr,
		sessionHeader, sessionHeader, sessionHeader,
		keyRing,
		ringClient,
	)

	// Advance the block height to the claim window open height.
	sharedParams := keepers.SharedKeeper.GetParams(ctx)
	claimMsgHeight := shared.GetClaimWindowOpenHeight(
		&sharedParams,
		sessionHeader.GetSessionEndBlockHeight(),
	)
	sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
	sdkCtx = sdkCtx.WithBlockHeight(claimMsgHeight)
	ctx = sdkCtx

	// Create a valid claim.
	keepers.CreateClaimAndStoreBlockHash(
		ctx, t, 1,
		supplierAddr,
		appAddr,
		service,
		sessionTree,
		sessionHeader,
	)

	proofWindowOpenHeight := shared.GetProofWindowOpenHeight(&sharedParams, sessionHeader.GetSessionEndBlockHeight())
	proofWindowCloseHeight := shared.GetProofWindowCloseHeight(&sharedParams, sessionHeader.GetSessionEndBlockHeight())

	tests := []struct {
		desc           string
		proofMsgHeight int64
		expectedErr    error
	}{
		{
			desc:           "proof message height equals proof window open height minus one",
			proofMsgHeight: int64(proofWindowOpenHeight) - 1,
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofProofOutsideOfWindow.Wrapf(
					"current block height (%d) is less than session proof window open height (%d)",
					int64(proofWindowOpenHeight)-1,
					proofWindowOpenHeight,
				).Error(),
			),
		},
		{
			desc:           "proof message height equals proof window close height plus one",
			proofMsgHeight: int64(proofWindowCloseHeight) + 1,
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofProofOutsideOfWindow.Wrapf(
					"current block height (%d) is greater than session proof window close height (%d)",
					int64(proofWindowCloseHeight)+1,
					proofWindowCloseHeight,
				).Error(),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// Advance the block height to the test proof msg height.
			sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
			sdkCtx = sdkCtx.WithBlockHeight(test.proofMsgHeight)
			ctx = sdkCtx

			proofMsg := common.NewTestProofMsg(t,
				supplierAddr,
				sessionHeader,
				sessionTree,
				expectedMerkleProofPath,
			)
			_, err := proofMsgServer.SubmitProof(ctx, proofMsg)
			require.ErrorContains(t, err, test.expectedErr.Error())

			proofRes, err := keepers.AllProofs(ctx, &types.QueryAllProofsRequest{})
			require.NoError(t, err)

			proofs := proofRes.GetProofs()
			require.Lenf(t, proofs, 0, "expected 0 proof, got %d", len(proofs))
		})
	}
}

func TestMsgServer_SubmitProof_Error(t *testing.T) {
	opts := []keepertest.ProofKeepersOpt{
		// Set block hash so we can have a deterministic expected on-chain proof requested by the protocol.
		common.WithBlockHash[
			keepertest.ProofKeepersOpt,
			*keepertest.ProofModuleKeepers,
		](blockHeaderHash),
		// Set block height to 1 so there is a valid session on-chain.
		common.WithBlockHeight[
			keepertest.ProofKeepersOpt,
			*keepertest.ProofModuleKeepers,
		](1),
	}
	keepers, ctx := keepertest.NewProofModuleKeepers(t, opts...)

	// Ensure the minimum relay difficulty bits is set to zero so that test cases
	// don't need to mine for valid relays.
	err := keepers.Keeper.SetParams(ctx, testProofParams)
	require.NoError(t, err)

	// Construct a keyring to hold the keypairs for the accounts used in the test.
	keyRing := keyring.NewInMemory(keepers.Codec)

	// The base session start height used for testing
	sessionStartHeight := int64(1)

	// Create a pre-generated account iterator to create accounts for the test.
	preGeneratedAccts := testkeyring.PreGeneratedAccounts()

	// Create accounts in the account keeper with corresponding keys in the keyring
	// for the applications and suppliers used in the tests.
	supplierAddr := testkeyring.CreateOnChainAccount(
		ctx, t,
		supplierUid,
		keyRing,
		keepers,
		preGeneratedAccts,
	).String()
	wrongSupplierAddr := testkeyring.CreateOnChainAccount(
		ctx, t,
		"wrong_supplier",
		keyRing,
		keepers,
		preGeneratedAccts,
	).String()
	appAddr := testkeyring.CreateOnChainAccount(
		ctx, t,
		"app",
		keyRing,
		keepers,
		preGeneratedAccts,
	).String()
	wrongAppAddr := testkeyring.CreateOnChainAccount(
		ctx, t,
		"wrong_app",
		keyRing,
		keepers,
		preGeneratedAccts,
	).String()

	service := &sharedtypes.Service{Id: testServiceId}
	wrongService := &sharedtypes.Service{Id: "wrong_svc"}

	// Add a supplier and application pair that are expected to be in the session.
	keepers.AddServiceActors(ctx, t, service, supplierAddr, appAddr)

	// Add a supplier and application pair that are *not* expected to be in the session.
	keepers.AddServiceActors(ctx, t, wrongService, wrongSupplierAddr, wrongAppAddr)

	// Get the session for the application/supplier pair which is expected
	// to be claimed and for which a valid proof would be accepted.
	validSessionHeader := common.GetSessionHeader(ctx, t, keepers, appAddr, service, 1)

	// Get the session for the application/supplier pair which is
	// *not* expected to be claimed.
	unclaimedSessionHeader := common.GetSessionHeader(ctx, t, keepers, wrongAppAddr, wrongService, 1)

	// Construct a session header with session ID that doesn't match the expected session ID.
	wrongSessionIdHeader := *validSessionHeader
	wrongSessionIdHeader.SessionId = "wrong session ID"

	// Construct a session header with a session start height that doesn't match
	// the expected session start height.
	wrongSessionStartHeightHeader := *validSessionHeader
	wrongSessionStartHeightHeader.SessionStartBlockHeight = 2

	// Construct a session header with a session end height that doesn't match
	// the expected session end height.
	wrongSessionEndHeightHeader := *validSessionHeader
	wrongSessionEndHeightHeader.SessionEndBlockHeight = 3

	// TODO_TECHDEBT: add a test case such that we can distinguish between early
	// & late session end block heights.

	// Construct a proof message server from the proof keeper.
	proofMsgServer := keeper.NewMsgServerImpl(*keepers.Keeper)

	// Construct a ringClient to get the application's ring & verify the relay
	// request signature.
	ringClient, err := rings.NewRingClient(depinject.Supply(
		polyzero.NewLogger(),
		types.NewAppKeeperQueryClient(keepers.ApplicationKeeper),
		types.NewAccountKeeperQueryClient(keepers.AccountKeeper),
		types.NewSharedKeeperQueryClient(keepers.SharedKeeper),
	))
	require.NoError(t, err)

	// Construct a valid session tree with 5 relays.
	numRelays := uint(5)
	validSessionTree := common.NewFilledSessionTree(
		ctx, t,
		numRelays,
		supplierUid, supplierAddr,
		validSessionHeader, validSessionHeader, validSessionHeader,
		keyRing,
		ringClient,
	)

	// Advance the block height to the claim window open height.
	sharedParams := keepers.SharedKeeper.GetParams(ctx)
	claimMsgHeight := shared.GetClaimWindowOpenHeight(
		&sharedParams,
		validSessionHeader.GetSessionEndBlockHeight(),
	)
	sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
	sdkCtx = sdkCtx.WithBlockHeight(claimMsgHeight)
	ctx = sdkCtx

	// Create a valid claim for the expected session and update the block hash
	// store for the corresponding session.
	keepers.CreateClaimAndStoreBlockHash(
		ctx, t, 1,
		supplierAddr,
		appAddr,
		service,
		validSessionTree,
		validSessionHeader,
	)

	// Compute the difficulty in bits of the closest relay from the valid session tree.
	validClosestRelayDifficultyBits := getClosestRelayDifficultyBits(t, validSessionTree, expectedMerkleProofPath)

	// Copy `emptyBlockHash` to `wrongClosestProofPath` to with a missing byte
	// so the closest proof is invalid (i.e. unmarshalable).
	invalidClosestProofBytes := make([]byte, len(expectedMerkleProofPath)-1)

	// Store the expected error returned during deserialization of the invalid
	// closest Merkle proof bytes.
	sparseMerkleClosestProof := &smt.SparseMerkleClosestProof{}
	expectedInvalidProofUnmarshalErr := sparseMerkleClosestProof.Unmarshal(invalidClosestProofBytes)

	// Construct a relay to be mangled such that it fails to deserialize in order
	// to set the error expectation for the relevant test case.
	mangledRelay := common.NewEmptyRelay(validSessionHeader, validSessionHeader)

	// Ensure valid relay request and response signatures.
	common.SignRelayRequest(ctx, t, mangledRelay, appAddr, keyRing, ringClient)
	common.SignRelayResponse(ctx, t, mangledRelay, supplierUid, supplierAddr, keyRing)

	// Serialize the relay so that it can be mangled.
	mangledRelayBz, err := mangledRelay.Marshal()
	require.NoError(t, err)

	// Mangle the serialized relay to cause an error during deserialization.
	// Mangling could involve any byte randomly being swapped to any value
	// so unmarshaling fails, but we are setting the first byte to 0 for simplicity.
	mangledRelayBz[0] = 0x00

	// Declare an invalid signature byte slice to construct expected relay request
	// and response errors and use in corresponding test cases.
	invalidSignatureBz := []byte("invalid signature bytes")

	// Prepare an invalid proof of the correct size.
	wrongClosestProofPath := make([]byte, len(expectedMerkleProofPath))
	copy(wrongClosestProofPath, expectedMerkleProofPath)
	copy(wrongClosestProofPath, "wrong closest proof path")

	// Increment the block height to the test proof height.
	proofMsgHeight := shared.GetProofWindowOpenHeight(
		&sharedParams,
		validSessionHeader.GetSessionEndBlockHeight(),
	)
	ctx = cosmostypes.UnwrapSDKContext(ctx).WithBlockHeight(proofMsgHeight)

	tests := []struct {
		desc        string
		newProofMsg func(t *testing.T) *types.MsgSubmitProof
		expectedErr error
	}{
		{
			desc: "proof service ID cannot be empty",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Set proof session ID to empty string.
				emptySessionIdHeader := *validSessionHeader
				emptySessionIdHeader.SessionId = ""

				// Construct new proof message.
				return common.NewTestProofMsg(t,
					supplierAddr,
					&emptySessionIdHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.InvalidArgument,
				types.ErrProofInvalidSessionId.Wrapf(
					"session ID does not match on-chain session ID; expected %q, got %q",
					validSessionHeader.GetSessionId(),
					"",
				).Error(),
			),
		},
		{
			desc: "merkle proof cannot be empty",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct new proof message.
				proof := common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)

				// Set merkle proof to an empty byte slice.
				proof.Proof = []byte{}
				return proof
			},
			expectedErr: status.Error(
				codes.InvalidArgument,
				types.ErrProofInvalidProof.Wrap(
					"proof cannot be empty",
				).Error(),
			),
		},
		{
			desc: "proof session ID must match on-chain session ID",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct new proof message using the wrong session ID.
				return common.NewTestProofMsg(t,
					supplierAddr,
					&wrongSessionIdHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.InvalidArgument,
				types.ErrProofInvalidSessionId.Wrapf(
					"session ID does not match on-chain session ID; expected %q, got %q",
					validSessionHeader.GetSessionId(),
					wrongSessionIdHeader.GetSessionId(),
				).Error(),
			),
		},
		{
			desc: "proof supplier must be in on-chain session",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct a proof message with a  supplier that does not belong in the session.
				return common.NewTestProofMsg(t,
					wrongSupplierAddr,
					validSessionHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.InvalidArgument,
				types.ErrProofNotFound.Wrapf(
					"supplier address %q not found in session ID %q",
					wrongSupplierAddr,
					validSessionHeader.GetSessionId(),
				).Error(),
			),
		},
		{
			desc: "merkle proof must be deserializable",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct new proof message.
				proof := common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)

				// Set merkle proof to an incorrect byte slice.
				proof.Proof = invalidClosestProofBytes

				return proof
			},
			expectedErr: status.Error(
				codes.InvalidArgument,
				types.ErrProofInvalidProof.Wrapf(
					"failed to unmarshal closest merkle proof: %s",
					expectedInvalidProofUnmarshalErr,
				).Error(),
			),
		},
		{
			desc: "relay must be deserializable",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct a session tree to which we'll add 1 unserializable relay.
				mangledRelaySessionTree := common.NewEmptySessionTree(t, validSessionHeader)

				// Add the mangled relay to the session tree.
				err = mangledRelaySessionTree.Update([]byte{1}, mangledRelayBz, 1)
				require.NoError(t, err)

				// Get the Merkle root for the session tree in order to construct a claim.
				mangledRelayMerkleRootBz, err := mangledRelaySessionTree.Flush()
				require.NoError(t, err)

				// Create a claim with a merkle root derived from a session tree
				// with an unserializable relay.
				claimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					mangledRelayMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, claimMsg)
				require.NoError(t, err)

				// Construct new proof message derived from a session tree
				// with an unserializable relay.
				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					mangledRelaySessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.InvalidArgument,
				types.ErrProofInvalidRelay.Wrapf(
					"failed to unmarshal relay: %s",
					keepers.Codec.Unmarshal(mangledRelayBz, &servicetypes.Relay{}),
				).Error(),
			),
		},
		{
			// TODO_TEST(community): expand: test case to cover each session header field.
			desc: "relay request session header must match proof session header",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct a session tree with 1 relay with a session header containing
				// a session ID that doesn't match the proof session ID.
				numRelays := uint(1)
				wrongRequestSessionIdSessionTree := common.NewFilledSessionTree(
					ctx, t,
					numRelays,
					supplierUid, supplierAddr,
					validSessionHeader, &wrongSessionIdHeader, validSessionHeader,
					keyRing,
					ringClient,
				)

				// Get the Merkle root for the session tree in order to construct a claim.
				wrongRequestSessionIdMerkleRootBz, err := wrongRequestSessionIdSessionTree.Flush()
				require.NoError(t, err)

				// Create a claim with a merkle root derived from a relay
				// request containing the wrong session ID.
				claimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					wrongRequestSessionIdMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, claimMsg)
				require.NoError(t, err)

				// Construct new proof message using the valid session header,
				// *not* the one used in the session tree's relay request.
				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					wrongRequestSessionIdSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofInvalidRelay.Wrapf(
					"session headers session IDs mismatch; expected: %q, got: %q",
					validSessionHeader.GetSessionId(),
					wrongSessionIdHeader.GetSessionId(),
				).Error(),
			),
		},
		{
			// TODO_TEST: expand: test case to cover each session header field.
			desc: "relay response session header must match proof session header",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct a session tree with 1 relay with a session header containing
				// a session ID that doesn't match the expected session ID.
				numRelays := uint(1)
				wrongResponseSessionIdSessionTree := common.NewFilledSessionTree(
					ctx, t,
					numRelays,
					supplierUid, supplierAddr,
					validSessionHeader, validSessionHeader, &wrongSessionIdHeader,
					keyRing,
					ringClient,
				)

				// Get the Merkle root for the session tree in order to construct a claim.
				wrongResponseSessionIdMerkleRootBz, err := wrongResponseSessionIdSessionTree.Flush()
				require.NoError(t, err)

				// Create a claim with a merkle root derived from a relay
				// response containing the wrong session ID.
				claimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					wrongResponseSessionIdMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, claimMsg)
				require.NoError(t, err)

				// Construct new proof message using the valid session header,
				// *not* the one used in the session tree's relay response.
				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					wrongResponseSessionIdSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofInvalidRelay.Wrapf(
					"session headers session IDs mismatch; expected: %q, got: %q",
					validSessionHeader.GetSessionId(),
					wrongSessionIdHeader.GetSessionId(),
				).Error(),
			),
		},
		{
			desc: "relay request signature must be valid",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Set the relay request signature to an invalid byte slice.
				invalidRequestSignatureRelay := common.NewEmptyRelay(validSessionHeader, validSessionHeader)
				invalidRequestSignatureRelay.Req.Meta.Signature = invalidSignatureBz

				// Ensure a valid relay response signature.
				common.SignRelayResponse(ctx, t, invalidRequestSignatureRelay, supplierUid, supplierAddr, keyRing)

				invalidRequestSignatureRelayBz, err := invalidRequestSignatureRelay.Marshal()
				require.NoError(t, err)

				// Construct a session tree with 1 relay with a session header containing
				// a session ID that doesn't match the expected session ID.
				invalidRequestSignatureSessionTree := common.NewEmptySessionTree(t, validSessionHeader)

				// Add the relay to the session tree.
				err = invalidRequestSignatureSessionTree.Update([]byte{1}, invalidRequestSignatureRelayBz, 1)
				require.NoError(t, err)

				// Get the Merkle root for the session tree in order to construct a claim.
				invalidRequestSignatureMerkleRootBz, err := invalidRequestSignatureSessionTree.Flush()
				require.NoError(t, err)

				// Create a claim with a merkle root derived from a session tree
				// with an invalid relay request signature.
				claimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					invalidRequestSignatureMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, claimMsg)
				require.NoError(t, err)

				// Construct new proof message derived from a session tree
				// with an invalid relay request signature.
				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					invalidRequestSignatureSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofInvalidRelayRequest.Wrapf(
					"error deserializing ring signature: %s",
					new(ring.RingSig).Deserialize(ring_secp256k1.NewCurve(), invalidSignatureBz),
				).Error(),
			),
		},
		{
			desc: "relay request signature is valid but signed by an incorrect application",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				t.Skip("TODO_TECHDEBT(@bryanchriswhite): Implement this")
				return nil
			},
		},
		{
			desc: "relay response signature must be valid",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Set the relay response signature to an invalid byte slice.
				relay := common.NewEmptyRelay(validSessionHeader, validSessionHeader)
				relay.Res.Meta.SupplierSignature = invalidSignatureBz

				// Ensure a valid relay request signature
				common.SignRelayRequest(ctx, t, relay, appAddr, keyRing, ringClient)

				relayBz, err := relay.Marshal()
				require.NoError(t, err)

				// Construct a session tree with 1 relay with a session header containing
				// a session ID that doesn't match the expected session ID.
				invalidResponseSignatureSessionTree := common.NewEmptySessionTree(t, validSessionHeader)

				// Add the relay to the session tree.
				err = invalidResponseSignatureSessionTree.Update([]byte{1}, relayBz, 1)
				require.NoError(t, err)

				// Get the Merkle root for the session tree in order to construct a claim.
				invalidResponseSignatureMerkleRootBz, err := invalidResponseSignatureSessionTree.Flush()
				require.NoError(t, err)

				// Create a claim with a merkle root derived from a session tree
				// with an invalid relay response signature.
				claimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					invalidResponseSignatureMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, claimMsg)
				require.NoError(t, err)

				// Construct new proof message derived from a session tree
				// with an invalid relay response signature.
				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					invalidResponseSignatureSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				servicetypes.ErrServiceInvalidRelayResponse.Wrap("invalid signature").Error(),
			),
		},
		{
			desc: "relay response signature is valid but signed by an incorrect supplier",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				t.Skip("TODO_TECHDEBT(@bryanchriswhite): Implement this")
				return nil
			},
		},
		{
			desc: "the merkle proof path provided does not match the one expected/enforced by the protocol",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct a new valid session tree for this test case because once the
				// closest proof has already been generated, the path cannot be changed.
				numRelays := uint(5)
				wrongPathSessionTree := common.NewFilledSessionTree(
					ctx, t,
					numRelays,
					supplierUid, supplierAddr,
					validSessionHeader, validSessionHeader, validSessionHeader,
					keyRing,
					ringClient,
				)

				wrongPathMerkleRootBz, err := wrongPathSessionTree.Flush()
				require.NoError(t, err)

				// Create a valid claim with the expected merkle root.
				claimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					wrongPathMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, claimMsg)
				require.NoError(t, err)

				// Construct new proof message derived from a session tree
				// with an invalid relay response signature.
				return common.NewTestProofMsg(t, supplierAddr, validSessionHeader, wrongPathSessionTree, wrongClosestProofPath)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofInvalidProof.Wrapf(
					"the proof for the path provided (%x) does not match one expected by the on-chain protocol (%x)",
					wrongClosestProofPath,
					blockHeaderHash,
				).Error(),
			),
		},
		{
			desc: "relay difficulty must be greater than or equal to minimum (zero difficulty)",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Set the minimum relay difficulty to a non-zero value such that the relays
				// constructed by the test helpers have a negligable chance of being valid.
				err := keepers.Keeper.SetParams(ctx, types.Params{
					MinRelayDifficultyBits: 10,
				})
				require.NoError(t, err)

				// Reset the minimum relay difficulty to zero after this test case.
				t.Cleanup(func() {
					err := keepers.Keeper.SetParams(ctx, types.DefaultParams())
					require.NoError(t, err)
				})

				// Construct a proof message with a session tree containing
				// a relay of insufficient difficulty.
				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofInvalidRelay.Wrapf(
					"relay difficulty %d is less than the minimum difficulty %d",
					validClosestRelayDifficultyBits,
					10,
				).Error(),
			),
		},
		{
			desc: "relay difficulty must be greater than or equal to minimum (non-zero difficulty)",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				t.Skip("TODO_TECHDEBT(@bryanchriswhite): Implement this")
				return nil
			},
		},
		{ // group: claim must exist for proof message
			desc: "claim must exist for proof message",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				// Construct a new session tree corresponding to the unclaimed session.
				numRelays := uint(5)
				unclaimedSessionTree := common.NewFilledSessionTree(
					ctx, t,
					numRelays,
					"wrong_supplier", wrongSupplierAddr,
					unclaimedSessionHeader, unclaimedSessionHeader, unclaimedSessionHeader,
					keyRing,
					ringClient,
				)

				// Discard session tree Merkle root because no claim is being created.
				// Session tree must be closed (flushed) to compute closest Merkle Proof.
				_, err = unclaimedSessionTree.Flush()
				require.NoError(t, err)

				// Construct new proof message using the supplier & session header
				// from the session which is *not* expected to be claimed.
				return common.NewTestProofMsg(t,
					wrongSupplierAddr,
					unclaimedSessionHeader,
					unclaimedSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofClaimNotFound.Wrapf(
					"no claim found for session ID %q and supplier %q",
					unclaimedSessionHeader.GetSessionId(),
					wrongSupplierAddr,
				).Error(),
			),
		},
		{
			desc: "Valid proof cannot validate claim with an incorrect root",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				numRelays := uint(10)
				wrongMerkleRootSessionTree := common.NewFilledSessionTree(
					ctx, t,
					numRelays,
					supplierUid, supplierAddr,
					validSessionHeader, validSessionHeader, validSessionHeader,
					keyRing,
					ringClient,
				)

				wrongMerkleRootBz, err := wrongMerkleRootSessionTree.Flush()
				require.NoError(t, err)

				// Reset the block height to the test claim msg height.
				sharedParams := keepers.SharedKeeper.GetParams(ctx)
				claimMsgHeight := shared.GetClaimWindowOpenHeight(
					&sharedParams,
					validSessionHeader.GetSessionEndBlockHeight(),
				)
				sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
				sdkCtx = sdkCtx.WithBlockHeight(claimMsgHeight)
				ctx = sdkCtx

				// Create a claim with the incorrect Merkle root.
				wrongMerkleRootClaimMsg := common.NewTestClaimMsg(t,
					sessionStartHeight,
					validSessionHeader.GetSessionId(),
					supplierAddr,
					appAddr,
					service,
					wrongMerkleRootBz,
				)
				_, err = proofMsgServer.CreateClaim(ctx, wrongMerkleRootClaimMsg)
				require.NoError(t, err)

				// Advance the block height to the test proof msg height.
				ctx = cosmostypes.UnwrapSDKContext(ctx).WithBlockHeight(proofMsgHeight)

				return common.NewTestProofMsg(t,
					supplierAddr,
					validSessionHeader,
					validSessionTree,
					expectedMerkleProofPath,
				)
			},
			expectedErr: status.Error(
				codes.FailedPrecondition,
				types.ErrProofInvalidProof.Wrap("invalid closest merkle proof").Error(),
			),
		},
		{
			desc: "claim and proof application addresses must match",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				t.Skip("this test case reduces to either the 'claim must exist for proof message' or 'proof session ID must match on-chain session ID cases")
				return nil
			},
		},
		{
			desc: "claim and proof service IDs must match",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				t.Skip("this test case reduces to either the 'claim must exist for proof message' or 'proof session ID must match on-chain session ID cases")
				return nil
			},
		},
		{
			desc: "claim and proof supplier addresses must match",
			newProofMsg: func(t *testing.T) *types.MsgSubmitProof {
				t.Skip("this test case reduces to either the 'claim must exist for proof message' or 'proof session ID must match on-chain session ID cases")
				return nil
			},
		},
	}

	// Submit the corresponding proof.
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			proofMsg := test.newProofMsg(t)
			submitProofRes, err := proofMsgServer.SubmitProof(ctx, proofMsg)

			require.ErrorContains(t, err, test.expectedErr.Error())
			require.Nil(t, submitProofRes)

			proofRes, err := keepers.AllProofs(ctx, &types.QueryAllProofsRequest{})
			require.NoError(t, err)

			// Expect zero proofs to have been persisted as all test cases are error cases.
			proofs := proofRes.GetProofs()
			require.Lenf(t, proofs, 0, "expected 0 proofs, got %d", len(proofs))
		})
	}
}

// getClosestRelayDifficultyBits returns the number of leading 0s (i.e. relay
// mining difficulty bits) in the relayHash stored in the sessionTree that is
// is closest to the merkle proof path provided.
func getClosestRelayDifficultyBits(
	t *testing.T,
	sessionTree relayer.SessionTree,
	closestMerkleProofPath []byte,
) uint64 {
	// Retrieve a merkle proof that is closest to the path provided
	closestMerkleProof, err := sessionTree.ProveClosest(closestMerkleProofPath)
	require.NoError(t, err)

	// Extract the Relay (containing the RelayResponse & RelayRequest) from the merkle proof.
	relay := new(servicetypes.Relay)
	relayBz := closestMerkleProof.GetValueHash(&keeper.SmtSpec)
	err = relay.Unmarshal(relayBz)
	require.NoError(t, err)

	// Retrieve the hash of the relay.
	relayHash, err := relay.GetHash()
	require.NoError(t, err)

	// Count the number of leading 0s in the relay hash to determine its difficulty.
	relayDifficultyBits, err := protocol.CountHashDifficultyBits(relayHash[:])
	require.NoError(t, err)

	return uint64(relayDifficultyBits)
}
