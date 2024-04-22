package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	tokenomicstypes "github.com/pokt-network/poktroll/x/tokenomics/types"
)

func TestMsgUpdateParam_UpdatesSingleParam(t *testing.T) {
	var expectedMinRelayDifficultyBits int64 = 8

	k, msgSrv, ctx := setupMsgServer(t)
	defaultParams := tokenomicstypes.DefaultParams()
	require.NoError(t, k.SetParams(ctx, defaultParams))

	require.NotEqual(t, uint64(expectedMinRelayDifficultyBits), defaultParams.ComputeUnitsToTokensMultiplier)

	updateParamMsg := &tokenomicstypes.MsgUpdateParam{
		Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		Name:      "compute_units_to_tokens_multiplier",
		AsType:    &tokenomicstypes.MsgUpdateParam_AsInt64{AsInt64: expectedMinRelayDifficultyBits},
	}

	res, err := msgSrv.UpdateParam(ctx, updateParamMsg)
	require.NoError(t, err)
	require.Equal(t, uint64(expectedMinRelayDifficultyBits), res.Params.ComputeUnitsToTokensMultiplier)

	// TODO_BLOCKER: once we have more than one param per module, add assertions
	// here which ensure that other params were not changed!
}
