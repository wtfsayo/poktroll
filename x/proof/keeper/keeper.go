package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/crypto"
	"github.com/pokt-network/poktroll/pkg/crypto/rings"
	"github.com/pokt-network/poktroll/pkg/polylog"
	_ "github.com/pokt-network/poktroll/pkg/polylog/polyzero"
	"github.com/pokt-network/poktroll/x/proof/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		sessionKeeper     types.SessionKeeper
		applicationKeeper types.ApplicationKeeper

		ringClient     crypto.RingClient
		accountQuerier client.AccountQueryClient
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	sessionKeeper types.SessionKeeper,
	applicationKeeper types.ApplicationKeeper,
	accountKeeper types.AccountKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	// TODO_TECHDEBT: Use cosmos-sdk based polylog implementation once available. Also remove polyzero import.
	polylogger := polylog.Ctx(context.Background())
	applicationQuerier := types.NewAppKeeperQueryClient(applicationKeeper)
	accountQuerier := types.NewAccountKeeperQueryClient(accountKeeper)

	// RingKeeperClient holds the logic of verifying RelayRequests ring signatures
	// for both on-chain and off-chain actors.
	// As it takes care of getting the ring signature details (i.e. application
	// and gateways pub keys) it uses an Application and Account querier interfaces
	// that are compatible with the environment it's being used in.
	// In this on-chain context, the Proof keeper supplies AppKeeperQueryClient and
	// AccountKeeperQueryClient that are thin wrappers around the Application and
	// Account keepers respectively, and satisfy the RingClient needs.
	ringKeeperClientDeps := depinject.Supply(polylogger, applicationQuerier, accountQuerier)
	ringKeeperClient, err := rings.NewRingClient(ringKeeperClientDeps)
	if err != nil {
		panic(err)
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		sessionKeeper:     sessionKeeper,
		applicationKeeper: applicationKeeper,

		ringClient:     ringKeeperClient,
		accountQuerier: accountQuerier,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
