package gateway_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/testutil/nullify"
	"github.com/pokt-network/poktroll/testutil/sample"
	gateway "github.com/pokt-network/poktroll/x/gateway/module"
	"github.com/pokt-network/poktroll/x/gateway/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GatewayList: []types.Gateway{
			{
				Address: sample.AccAddress(),
				Stake:   sdk.NewCoin("upokt", math.NewInt(1000)),
			},
			{
				Address: sample.AccAddress(),
				Stake:   sdk.NewCoin("upokt", math.NewInt(1000)),
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GatewayKeeper(t)
	gateway.InitGenesis(ctx, k, genesisState)
	got := gateway.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GatewayList, got.GatewayList)
	// this line is used by starport scaffolding # genesis/test/assert
}
