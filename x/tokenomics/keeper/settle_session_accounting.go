package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	suppliertypes "github.com/pokt-network/poktroll/x/supplier/types"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
	"github.com/pokt-network/smt"
)

// SettleSessionAccounting implements TokenomicsKeeper#SettleSessionAccounting
func (k TokenomicsKeeper) SettleSessionAccounting(
	goCtx context.Context,
	claim *suppliertypes.Claim,
) error {
	if claim == nil {
		return types.ErrTokenomicsClaimNil
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx).With("method", "SettleSessionAccounting")

	// Decompose the claim into its constituent parts
	supplierAddress := sdk.AccAddress(claim.SupplierAddress)
	applicationAddress := sdk.AccAddress(claim.SessionHeader.ApplicationAddress)
	sessionHeader := claim.SessionHeader
	root := (smt.MerkleRoot)(claim.RootHash)

	// Make sure the session header is not nil
	if sessionHeader == nil {
		return types.ErrTokenomicsSessionHeaderNil
	}

	// Retrieve the existing tokenomics params
	params := k.GetParams(ctx)

	// Retrieve the sum of the root as a proxy into the amount of work done
	computeUnits := root.Sum()

	// Calculate the amount of tokens to mint & burn
	// TODO_BLOCKER: This is just a simple naive implementation
	upokt := sdk.NewInt(int64(computeUnits * params.ComputeUnitsToTokensMultiplier))
	upoktCoins := sdk.NewCoins(sdk.NewCoin("upokt", upokt))

	// NB: We are doing a mint & burn + transfer, instead of a simple transfer
	// of funds from the supplier to the application in order to enable second
	// order economic effects with more optionality.

	// Mint uPOKT to the supplier module account
	if err := k.bankKeeper.MintCoins(ctx, suppliertypes.ModuleName, upoktCoins); err != nil {
		return types.ErrTokenomicsApplicationModuleFeeFailed
	}

	logger.Info(fmt.Sprintf("minted %d uPOKT in the supplier module", upokt))

	// Sent the minted coins to the supplier
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, suppliertypes.ModuleName, supplierAddress, upoktCoins,
	); err != nil {
		return types.ErrTokenomicsApplicationModuleFeeFailed
	}

	logger.Info(fmt.Sprintf("sent %d uPOKT to supplier with address %s", upokt, supplierAddress))

	// Send coins from the application to the application module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, applicationAddress, apptypes.ModuleName, upoktCoins,
	); err != nil {
		return types.ErrTokenomicsApplicationModuleFeeFailed
	}

	logger.Info(fmt.Sprintf("took %d uPOKT from application with address %s", upokt, applicationAddress))

	// Burn uPOKT from the application module account
	if err := k.bankKeeper.BurnCoins(ctx, apptypes.ModuleName, upoktCoins); err != nil {
		return types.ErrTokenomicsApplicationModuleBurn
	}

	logger.Info(fmt.Sprintf("burned %d uPOKT in the application module", upokt))

	return nil
}