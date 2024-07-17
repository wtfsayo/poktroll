package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/pokt-network/poktroll/proto/types/tokenomics"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params tokenomics.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	paramsBz := store.Get(types.ParamsKey)
	if paramsBz == nil {
		return params
	}

	k.cdc.MustUnmarshal(paramsBz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params tokenomics.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	paramsBz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, paramsBz)

	return nil
}
