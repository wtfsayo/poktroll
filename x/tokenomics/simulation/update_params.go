package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/pokt-network/poktroll/x/tokenomics/keeper"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
)

func SimulateMsgUpdateParams(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.TokenomicsKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateParams{
			Authority: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateParams simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateParams simulation not implemented"), nil, nil
	}
}
