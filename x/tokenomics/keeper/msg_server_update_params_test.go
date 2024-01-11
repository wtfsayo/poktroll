package keeper_test

import (
	"testing"

	testkeeper "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/x/tokenomics/keeper"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
	"github.com/stretchr/testify/require"
)

func TestUpdateParams_Validity(t *testing.T) {
	tokenomicsKeeper, ctx := testkeeper.TokenomicsKeeper(t)
	srv := keeper.NewMsgServerImpl(*tokenomicsKeeper)

	params := types.DefaultParams()
	tokenomicsKeeper.SetParams(ctx, params)

	tests := []struct {
		desc string

		req *types.MsgUpdateParams

		expectErr     bool
		expectedPanic bool
		expErrMsg     string
	}{
		{
			desc: "invalid authority address",

			req: &types.MsgUpdateParams{
				Authority: "invalid",
			},

			expectErr:     true,
			expectedPanic: false,
			expErrMsg:     "invalid authority",
		},
		{
			desc: "invalid ComputeUnitsToTokensMultiplier",

			req: &types.MsgUpdateParams{
				Authority: tokenomicsKeeper.GetAuthority(),

				Params: types.Params{
					ComputeUnitsToTokensMultiplier: 0,
				},
			},

			expectErr:     true,
			expectedPanic: true,
			expErrMsg:     "invalid compute to tokens multiplier",
		},
		{
			desc: "successful param update",

			req: &types.MsgUpdateParams{
				Authority: tokenomicsKeeper.GetAuthority(),

				Params: types.Params{
					ComputeUnitsToTokensMultiplier: 1000000,
				},
			},

			expectedPanic: false,
			expectErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.expectedPanic {
				defer func() {
					if r := recover(); r != nil {
						_, err := srv.UpdateParams(ctx, tt.req)
						require.Error(t, err)
					}
				}()
				return
			}
			_, err := srv.UpdateParams(ctx, tt.req)
			if tt.expectErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.expErrMsg)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestUpdateParams_ComputeUnitsToTokensMultiplier(t *testing.T) {
	tokenomicsKeeper, ctx := testkeeper.TokenomicsKeeper(t)
	// srv := keeper.NewMsgServerImpl(*tokenomicsKeeper)

	params := types.DefaultParams()
	tokenomicsKeeper.SetParams(ctx, params)

	getParamsReq := &types.QueryParamsRequest{}
	getParamsRes, err := tokenomicsKeeper.Params(ctx, getParamsReq)
	require.Nil(t, err)
	require.Equal(t, 42, getParamsRes.Params.ComputeUnitsToTokensMultiplier)

	// req := &types.MsgUpdateParams{
	// 	Authority: tokenomicsKeeper.GetAuthority(),

	// 	Params: types.Params{
	// 		ComputeUnitsToTokensMultiplier: 1000000,
	// 	},
	// }

	// srv.GetParams(ctx, &types.QueryParamsRequest{})
}