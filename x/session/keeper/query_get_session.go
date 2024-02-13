package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pokt-network/poktroll/x/session/types"
)

// GetSession should be deterministic and always return the same session for
// the same block height.
func (k Keeper) GetSession(goCtx context.Context, req *types.QueryGetSessionRequest) (*types.QueryGetSessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if err := req.ValidateBasic(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Note that `GetSession` is called via the `Query` service rather than the `Msg` server.
	// The former is stateful but does not lead to state transitions, while the latter one
	// does. The request height depends on how much the node has synched and only acts as a read,
	// while the `Msg` server handles the code flow of the validator/sequencer when a new block
	// is being proposed.
	blockHeight := req.BlockHeight

	k.Logger(ctx).Info("Getting session for height: %d", blockHeight)

	sessionHydrator := NewSessionHydrator(req.ApplicationAddress, req.Service.Id, blockHeight)
	session, err := k.HydrateSession(ctx, sessionHydrator)
	if err != nil {
		return nil, err
	}

	res := &types.QueryGetSessionResponse{
		Session: session,
	}
	return res, nil
}
