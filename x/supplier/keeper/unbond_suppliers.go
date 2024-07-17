package keeper

import (
	"context"
	"fmt"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pokt-network/poktroll/x/shared"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
	"github.com/pokt-network/poktroll/x/supplier/types"
)

// EndBlockerUnbondSupplier unbonds suppliers that have finished the unbonding period.
func (k Keeper) EndBlockerUnbondSupplier(ctx context.Context) error {
	sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
	currentHeight := sdkCtx.BlockHeight()

	logger := k.Logger().With("method", "UnbondSupplier")

	// Iterate over all suppliers and unbond suppliers that have finished the unbonding period.
	// TODO_IMPROVE: Use an index to iterate over suppliers that have initiated the
	// unbonding action instead of iterating over all suppliers.
	for _, supplier := range k.GetAllSuppliers(ctx) {
		// Ignore suppliers that have not initiated the unbonding action.
		if supplier.UnstakeCommitSessionEndHeight == types.SupplierNotUnstaking {
			continue
		}

		unbondingHeight := k.GetSupplierUnbondingHeight(ctx, &supplier)

		if unbondingHeight <= currentHeight {

			// Retrieve the address of the supplier.
			supplierAddress, err := cosmostypes.AccAddressFromBech32(supplier.Address)
			if err != nil {
				logger.Error(fmt.Sprintf("could not parse address %s", supplier.Address))
				return err
			}

			// Send the coins from the supplier pool back to the supplier.
			if err = k.bankKeeper.SendCoinsFromModuleToAccount(
				ctx, types.ModuleName, supplierAddress, []cosmostypes.Coin{*supplier.Stake},
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

// GetSupplierUnbondingHeight returns the height at which the supplier can be unbonded.
func (k Keeper) GetSupplierUnbondingHeight(
	ctx context.Context,
	supplier *sharedtypes.Supplier,
) int64 {
	sharedParams := k.sharedKeeper.GetParams(ctx)

	return shared.GetProofWindowCloseHeight(
		&sharedParams,
		int64(supplier.UnstakeCommitSessionEndHeight),
	)
}
