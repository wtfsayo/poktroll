package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/pokt-network/poktroll/x/gateway/keeper"
	"github.com/pokt-network/poktroll/x/gateway/types"
)

// SimulateMsgUnstakeGateway simulates the unstake gateway operation.
// TODO_TECHDEBT: Implement simulation logic.
func SimulateMsgUnstakeGateway(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUnstakeGateway{
			Address: simAccount.Address.String(),
		}

		// TODO: Handling the UnstakeGateway simulation

		return simtypes.NoOpMsg(
			types.ModuleName,
			sdk.MsgTypeURL(msg),
			"UnstakeGateway simulation not implemented",
		), nil, nil
	}
}
