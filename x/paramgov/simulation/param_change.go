package simulation

import (
	"math/rand"

	"github.com/pokt-network/poktroll/x/paramgov/keeper"
	"github.com/pokt-network/poktroll/x/paramgov/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgParamChange(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgParamChange{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ParamChange simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ParamChange simulation not implemented"), nil, nil
	}
}
