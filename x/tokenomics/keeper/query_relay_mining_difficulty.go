package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pokt-network/poktroll/proto/types/tokenomics"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
)

func (k Keeper) RelayMiningDifficultyAll(ctx context.Context, req *tokenomics.QueryAllRelayMiningDifficultyRequest) (*tokenomics.QueryAllRelayMiningDifficultyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var relayMiningDifficulties []tokenomics.RelayMiningDifficulty

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	relayMiningDifficultyStore := prefix.NewStore(store, types.KeyPrefix(types.RelayMiningDifficultyKeyPrefix))

	pageRes, err := query.Paginate(relayMiningDifficultyStore, req.Pagination, func(key []byte, value []byte) error {
		var relayMiningDifficulty tokenomics.RelayMiningDifficulty
		if err := k.cdc.Unmarshal(value, &relayMiningDifficulty); err != nil {
			return err
		}

		relayMiningDifficulties = append(relayMiningDifficulties, relayMiningDifficulty)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tokenomics.QueryAllRelayMiningDifficultyResponse{RelayMiningDifficulty: relayMiningDifficulties, Pagination: pageRes}, nil
}

func (k Keeper) RelayMiningDifficulty(ctx context.Context, req *tokenomics.QueryGetRelayMiningDifficultyRequest) (*tokenomics.QueryGetRelayMiningDifficultyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	difficulty, found := k.GetRelayMiningDifficulty(
		ctx,
		req.ServiceId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &tokenomics.QueryGetRelayMiningDifficultyResponse{RelayMiningDifficulty: difficulty}, nil
}
