//go:generate mockgen -destination ../../../testutil/tokenomics/mocks/expected_keepers_mock.go -package mocks . AccountKeeper,BankKeeper,ApplicationKeeper,ProofKeeper

package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	apptypes "github.com/pokt-network/poktroll/x/application/types"
	prooftypes "github.com/pokt-network/poktroll/x/proof/types"
	sessiontypes "github.com/pokt-network/poktroll/x/session/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI // only used for simulation
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	// We use the bankkeeper SendXXX instead of DelegateXX methods
	// because their purpose is to "escrow" funds on behalf of an account rather
	// than "delegate" funds from one account to another which is more closely
	// linked to staking.
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type ApplicationKeeper interface {
	GetApplication(ctx context.Context, address string) (app apptypes.Application, found bool)
	GetAllApplications(ctx context.Context) []apptypes.Application
	SetApplication(context.Context, apptypes.Application)
}

type ProofKeeper interface {
	GetAllClaims(ctx context.Context) []prooftypes.Claim
	RemoveClaim(ctx context.Context, sessionId, supplierAddr string)
	GetProof(ctx context.Context, sessionId, supplierAddr string) (proof prooftypes.Proof, isProofFound bool)
	RemoveProof(ctx context.Context, sessionId, supplierAddr string)

	// Only used for testing & simulation
	UpsertClaim(ctx context.Context, claim prooftypes.Claim)
	UpsertProof(ctx context.Context, claim prooftypes.Proof)

	SetParams(ctx context.Context, params prooftypes.Params) error
}

type SharedKeeper interface {
	GetParams(ctx context.Context) sharedtypes.Params
}

type SupplierKeeper interface {
	SetSupplier(context.Context, sharedtypes.Supplier)
}

type SessionKeeper interface {
	GetSession(context.Context, *sessiontypes.QueryGetSessionRequest) (*sessiontypes.QueryGetSessionResponse, error)
	StoreBlockHash(ctx context.Context)
}
