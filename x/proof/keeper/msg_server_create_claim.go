package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pokt-network/poktroll/telemetry"
	"github.com/pokt-network/poktroll/x/proof/types"
)

func (k msgServer) CreateClaim(ctx context.Context, msg *types.MsgCreateClaim) (*types.MsgCreateClaimResponse, error) {
	// TODO_BLOCKER(@bryanchriswhite): Prevent Claim upserts after the ClaimWindow is closed.
	// TODO_BLOCKER(@bryanchriswhite): Validate the signature on the Claim message corresponds to the supplier before Upserting.

	isSuccessful := false
	defer telemetry.EventSuccessCounter(
		"create_claim",
		telemetry.DefaultCounterFn,
		func() bool { return isSuccessful },
	)

	logger := k.Logger().With("method", "CreateClaim")
	logger.Info("creating claim")

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	sessionHeader := msg.GetSessionHeader()
	session, err := k.queryAndValidateSessionHeader(
		ctx,
		sessionHeader,
		msg.GetSupplierAddress(),
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger = logger.
		With(
			"session_id", session.GetSessionId(),
			"session_end_height", sessionHeader.GetSessionEndBlockHeight(),
			"supplier", msg.GetSupplierAddress(),
		)

	/*
		TODO_BLOCKER(@bryanchriswhite):

		### Msg distribution validation (depends on sessionRes validation)
		1. [ ] governance-based earliest block offset
		2. [ ] pseudo-randomize earliest block offset

		### Claim validation
		1. [x] sessionRes validation
		2. [ ] msg distribution validation
	*/

	logger.Info("validated claim")

	// Construct and upsert claim after all validation.
	claim := types.Claim{
		SupplierAddress: msg.GetSupplierAddress(),
		SessionHeader:   sessionHeader,
		RootHash:        msg.GetRootHash(),
	}

	// TODO_BLOCKER(@Olshansk): check if this claim already exists and return an
	// appropriate error in any case where the supplier should no longer be able
	// to update the given proof.
	k.Keeper.UpsertClaim(ctx, claim)

	logger.Info("created new claim")

	isSuccessful = true
	// TODO_BETA: return the claim in the response.
	return &types.MsgCreateClaimResponse{}, nil
}
