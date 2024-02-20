package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/cmd/poktrolld/cmd"
	keepertest "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/testutil/sample"
	"github.com/pokt-network/poktroll/x/session/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

func init() {
	cmd.InitSDKConfig()
}

// NOTE: See `session_hydrator_test.go` for more extensive test coverage of different
// GetSession scenarios. This is just used to verify a few basic scenarios that act as
// the Cosmos SDK context aware wrapper around it.

func TestSession_GetSession_Success(t *testing.T) {
	keeper, ctx := keepertest.SessionKeeper(t)
	ctx = sdk.UnwrapSDKContext(ctx).WithBlockHeight(100) // provide a sufficiently large block height to avoid errors

	// TODO_TECHDEBT(#377): These test assume that the genesis block has a height of 0,
	// rewrite them in terms of height = 1 genesis.
	type test struct {
		name string

		appAddr     string
		serviceId   string
		blockHeight int64

		expectedSessionId     string
		expectedSessionNumber int64
		expectedNumSuppliers  int
	}

	tests := []test{
		{
			name: "valid - app1 svc1 at height=1",

			appAddr:     keepertest.TestApp1Address,
			serviceId:   keepertest.TestServiceId1,
			blockHeight: 1,

			// Intentionally only checking a subset of the session metadata returned
			expectedSessionId:     "6f2e0b6cba5a8cb93506ed4045143c4268945ebfb730b2c98fc7e3dc40132926",
			expectedSessionNumber: 0,
			expectedNumSuppliers:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &types.QueryGetSessionRequest{
				ApplicationAddress: test.appAddr,
				Service: &sharedtypes.Service{
					Id: test.serviceId,
				},
				BlockHeight: 1,
			}

			response, err := keeper.GetSession(ctx, req)
			require.NoError(t, err)
			require.NotNil(t, response)

			require.Equal(t, test.expectedSessionId, response.Session.SessionId)
			require.Equal(t, test.expectedSessionNumber, response.Session.SessionNumber)
			require.Len(t, response.Session.Suppliers, test.expectedNumSuppliers)
		})
	}
}

func TestSession_GetSession_Failure(t *testing.T) {
	keeper, ctx := keepertest.SessionKeeper(t)
	ctx = sdk.UnwrapSDKContext(ctx).WithBlockHeight(100) // provide a sufficiently large block height to avoid errors

	type test struct {
		name string

		appAddr     string
		serviceId   string
		blockHeight int64

		expectedErrContains string
	}

	tests := []test{
		{
			name: "application address does not reflected a staked application",

			appAddr:     sample.AccAddress(), // a random (valid) app address that's not staked
			serviceId:   keepertest.TestServiceId1,
			blockHeight: 1,

			expectedErrContains: types.ErrSessionAppNotFound.Error(),
		},
		{
			name: "application staked for service that has no available suppliers",

			appAddr:     keepertest.TestApp1Address,
			serviceId:   keepertest.TestServiceId11,
			blockHeight: 1,

			expectedErrContains: types.ErrSessionSuppliersNotFound.Error(),
		},
		{
			name: "application is valid but not staked for the specified service",

			appAddr:     keepertest.TestApp1Address,
			serviceId:   "svc9001", // App1 is not staked for service over 9000
			blockHeight: 1,

			expectedErrContains: types.ErrSessionAppNotStakedForService.Error(),
		},
		{
			name: "application address is invalid format",

			appAddr:     "invalid_app_address",
			serviceId:   keepertest.TestServiceId1,
			blockHeight: 1,

			expectedErrContains: types.ErrSessionInvalidAppAddress.Error(),
		},
		{
			name: "service ID is invalid",

			appAddr:     keepertest.TestApp1Address,
			serviceId:   "service_id_is_too_long_to_be_valid",
			blockHeight: 1,

			expectedErrContains: "invalid service in session",
		},
		{
			name: "negative block height",

			appAddr:     keepertest.TestApp1Address,
			serviceId:   keepertest.TestServiceId1,
			blockHeight: -1,

			expectedErrContains: "invalid block height for session being retrieved",
		},
	}

	expectedRes := (*types.QueryGetSessionResponse)(nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &types.QueryGetSessionRequest{
				ApplicationAddress: test.appAddr,
				Service: &sharedtypes.Service{
					Id: test.serviceId,
				},
				BlockHeight: test.blockHeight,
			}

			res, err := keeper.GetSession(ctx, req)
			require.Error(t, err)
			require.Contains(t, err.Error(), test.expectedErrContains)
			require.Equal(t, expectedRes, res)
		})
	}
}
