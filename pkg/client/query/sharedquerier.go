package query

import (
	"context"

	"cosmossdk.io/depinject"
	"github.com/cosmos/gogoproto/grpc"

	"github.com/pokt-network/poktroll/pkg/client"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

var _ client.SharedQueryClient = (*sharedQuerier)(nil)

// sharedQuerier is a wrapper around the sharedtypes.QueryClient that enables the
// querying of on-chain shared information through a single exposed method
// which returns an sharedtypes.Session struct
type sharedQuerier struct {
	clientConn    grpc.ClientConn
	sharedQuerier sharedtypes.QueryClient
}

// NewSharedQuerier returns a new instance of a client.SharedQueryClient by
// injecting the dependecies provided by the depinject.Config.
//
// Required dependencies:
// - clientCtx
func NewSharedQuerier(deps depinject.Config) (client.SharedQueryClient, error) {
	querier := &sharedQuerier{}

	if err := depinject.Inject(
		deps,
		&querier.clientConn,
	); err != nil {
		return nil, err
	}

	querier.sharedQuerier = sharedtypes.NewQueryClient(querier.clientConn)

	return querier, nil
}

// GetParams queries & returns the shared module on-chain parameters.
//
// TODO_TECHDEBT(#543): We don't really want to have to query the params for every method call.
// Once `ModuleParamsClient` is implemented, use its replay observable's `#Last` method
// to get the most recently (asynchronously) observed (and cached) value.
func (sessq *sharedQuerier) GetParams(ctx context.Context) (*sharedtypes.Params, error) {
	req := &sharedtypes.QueryParamsRequest{}
	res, err := sessq.sharedQuerier.Params(ctx, req)
	if err != nil {
		return nil, ErrQuerySessionParams.Wrapf("[%v]", err)
	}
	return &res.Params, nil
}