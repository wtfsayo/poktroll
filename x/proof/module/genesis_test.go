package proof_test

import (
	"testing"

	keepertest "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/testutil/nullify"
	"github.com/pokt-network/poktroll/testutil/sample"
	"github.com/pokt-network/poktroll/x/proof/module"
	"github.com/pokt-network/poktroll/x/proof/types"
	sessiontypes "github.com/pokt-network/poktroll/x/session/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	mockSessionId := "mock_session_id"

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimList: []types.Claim{
			{
				SupplierAddress: sample.AccAddress(),
				SessionHeader: &sessiontypes.SessionHeader{
					SessionId:          mockSessionId,
					ApplicationAddress: sample.AccAddress(),
				},
				RootHash: []byte{1, 2, 3},
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProofKeeper(t)
	proof.InitGenesis(ctx, k, genesisState)
	got := proof.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClaimList, got.ClaimList)
	// this line is used by starport scaffolding # genesis/test/assert
}
