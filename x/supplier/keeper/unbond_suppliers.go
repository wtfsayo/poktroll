package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pokt-network/poktroll/x/shared"
	"github.com/pokt-network/poktroll/x/supplier/types"
)

// EndBlockerUnbondSupplier unbonds suppliers that have finished the unbonding period.
func (k Keeper) EndBlockerUnbondSupplier(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := sdkCtx.BlockHeight()
	sharedParams := k.sharedKeeper.GetParams(ctx)
	sessionEndHeight := shared.GetSessionEndHeight(&sharedParams, currentHeight)

	// Only process unbonding at the end of the session.
	if currentHeight != sessionEndHeight {
		return nil
	}

	logger := k.Logger().With("method", "UnbondSupplier")

	// Iterate over all suppliers and unbond suppliers that have finished the unbonding period.
	// TODO_IMPROVE: Use an index to iterate over suppliers that have initiated the
	// unbonding action instead of iterating over all suppliers.
	for _, supplier := range k.GetAllSuppliers(ctx) {
		// Ignore suppliers that have not initiated the unbonding action.
		if supplier.UnstakeCommitSessionEndHeight == 0 {
			continue
		}

		unbondingPeriodBlocks := k.GetParams(ctx).SupplierUnbondingPeriodBlocks
		unbondingHeight := supplier.UnstakeCommitSessionEndHeight + unbondingPeriodBlocks

		if unbondingHeight <= uint64(currentHeight) {

			// Retrieve the address of the supplier.
			supplierAddress, err := sdk.AccAddressFromBech32(supplier.Address)
			if err != nil {
				logger.Error(fmt.Sprintf("could not parse address %s", supplier.Address))
				return err
			}

			// Send the coins from the supplier pool back to the supplier.
			if err = k.bankKeeper.SendCoinsFromModuleToAccount(
				ctx, types.ModuleName, supplierAddress, []sdk.Coin{*supplier.Stake},
			); err != nil {
				logger.Error(fmt.Sprintf(
					"could not send %v coins from %s module to %s account due to %v",
					supplier.Stake, supplierAddress, types.ModuleName, err,
				))
				return err
			}

			// Remove the supplier from the store.
			k.RemoveSupplier(ctx, supplierAddress.String())
			logger.Info(fmt.Sprintf("Successfully removed the supplier: %+v", supplier))
		}
	}

	return nil
}