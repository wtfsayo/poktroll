package basenet

import (
	"encoding/json"
	"fmt"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	testcli "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/testutil/network"
	"github.com/pokt-network/poktroll/testutil/testkeyring"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	gatewaytypes "github.com/pokt-network/poktroll/x/gateway/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
	suppliertypes "github.com/pokt-network/poktroll/x/supplier/types"
)

const warnNoModuleGenesisFmt = "WARN: no %s module genesis state found, if this is unexpected, ensure that genesis is populated before creating on-chain accounts"

// CreateKeyringAccounts populates the Keyring associated with the in-memory
// network with memnet.numKeyringAccounts() number of pre-generated accounts.
func (memnet *BaseInMemoryCosmosNetwork) CreateKeyringAccounts(t *testing.T) {
	t.Helper()

	// Keyring MAY be provided setting InMemoryNetworkConfig#Keyring.
	if memnet.Config.Keyring == nil {
		t.Log("Keyring not initialized, using new in-memory keyring")

		// Construct an in-memory keyring so that it can be populated and used prior
		// to network start.
		memnet.Config.Keyring = keyring.NewInMemory(memnet.Config.CosmosCfg.Codec)
	} else {
		t.Log("Keyring already initialized, using existing keyring")
	}

	// Create memnet.NumKeyringAccounts() number of accounts in the configured keyring.
	accts := testkeyring.CreatePreGeneratedKeyringAccounts(
		t, memnet.Config.Keyring, memnet.Config.GetNumKeyringAccounts(t),
	)

	// Assign the memnet's pre-generated accounts to a new pre-generated
	// accounts iterator containing only the accounts which were also created
	// in the keyring.
	memnet.PreGeneratedAccountIterator = testkeyring.NewPreGeneratedAccountIterator(accts...)
}

// CreateOnChainAccounts creates on-chain accounts (i.e. auth module) for the sum of
// the configured number of suppliers, applications, and gateways.
func (memnet *BaseInMemoryCosmosNetwork) CreateOnChainAccounts(t *testing.T) {
	t.Helper()

	// NB: while it may initially seem like the memnet#Init<actor>AccountsWithSequence() methods
	// can be refactored into a generic function, this is not possible so long as the genesis
	// state lists are passed directly & remain a slice of concrete types (as opposed to pointers).
	// Under these conditions, a generic function would not be able to unmarshal the genesis state
	// stored in the in-memory network because it is unmarshalling uses reflection, and it is not
	// possible to reflect over a nil generic type value.

	// Retrieve the supplier module's genesis state from cosmos-sdk in-memory network.
	supplierGenesisState := network.GetGenesisState[*suppliertypes.GenesisState](t, suppliertypes.ModuleName, memnet)
	if supplierGenesisState == nil {
		t.Logf(warnNoModuleGenesisFmt, "supplier")
	} else {
		// Initialize on-chain accounts for genesis suppliers.
		memnet.InitSupplierAccountsWithSequence(t, supplierGenesisState.SupplierList...)

	}

	// Retrieve the application module's genesis state from cosmos-sdk in-memory network.
	appGenesisState := network.GetGenesisState[*apptypes.GenesisState](t, apptypes.ModuleName, memnet)
	if appGenesisState == nil {
		t.Logf(warnNoModuleGenesisFmt, "application")
	} else {
		// Initialize on-chain accounts for genesis applications.
		memnet.InitAppAccountsWithSequence(t, appGenesisState.ApplicationList...)
	}

	// Retrieve the gateway module's genesis state from cosmos-sdk in-memory network.
	gatewayGenesisState := network.GetGenesisState[*gatewaytypes.GenesisState](t, gatewaytypes.ModuleName, memnet)
	if gatewayGenesisState == nil {
		t.Logf(warnNoModuleGenesisFmt, "gateway")
	} else {
		// Initialize on-chain accounts for genesis gateways.
		memnet.InitGatewayAccountsWithSequence(t, gatewayGenesisState.GatewayList...)
	}

	// need to wait for the account to be initialized in the next block
	require.NoError(t, memnet.GetNetwork(t).WaitForNextBlock())
}

// InitAccount initializes an Account by sending it some funds from the validator
// in the network to the address provided
func (memnet *BaseInMemoryCosmosNetwork) InitAccount(
	t *testing.T,
	addr types.AccAddress,
) {
	t.Helper()

	signerAccountNumber := 0
	// Validator's client context MUST be used for this CLI command because its keyring includes the validator's key
	clientCtx := memnet.Network.Validators[0].ClientCtx
	// MUST NOT use memnet.GetClientCtx(t) as its keyring does not include the validator's account
	// TODO_UPNEXT(@bryanchriswhite): Ensure validator key is always available via the in-memory network's keyring.
	net := memnet.GetNetwork(t)
	val := net.Validators[0]

	args := []string{
		fmt.Sprintf("--%s=true", flags.FlagOffline),
		fmt.Sprintf("--%s=%d", flags.FlagAccountNumber, signerAccountNumber),
		fmt.Sprintf("--%s=%d", flags.FlagSequence, memnet.NextAccountSequenceNumber(t)),

		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types.NewCoins(types.NewCoin(net.Config.BondDenom, math.NewInt(10))).String()),
	}
	amount := types.NewCoins(types.NewCoin("stake", math.NewInt(200)))
	responseRaw, err := testcli.MsgSendExec(clientCtx, val.Address, addr, amount, args...)
	require.NoError(t, err)
	var responseJSON map[string]interface{}
	err = json.Unmarshal(responseRaw.Bytes(), &responseJSON)
	require.NoError(t, err)
	require.Equal(t, float64(0), responseJSON["code"], "code is not 0 in the response: %v", responseJSON)
}

// GetPreGeneratedAccountIterator returns the pre-generated account iterator associated
// with the in-memory network; i.e. used to populate genesis and the keyring. Calling
// #Next() will return the next (unused) pre-generated account in the iterator.
func (memnet *BaseInMemoryCosmosNetwork) CreateNewOnChainAccount(t *testing.T) *testkeyring.PreGeneratedAccount {
	t.Helper()

	preGeneratedAcct, ok := testkeyring.PreGeneratedAccounts().Next()
	require.Truef(t, ok, "no pre-generated accounts available")

	// Create a supplier on-chain account.
	memnet.InitAccount(t, preGeneratedAcct.Address)

	testkeyring.CreatePreGeneratedKeyringAccounts(t, memnet.GetClientCtx(t).Keyring, 1)

	return preGeneratedAcct
}

// InitSupplierAccountsWithSequence uses the testCLI to create on-chain accounts
// (i.e. auth module) for the addresses of the given suppliers.
func (memnet *BaseInMemoryCosmosNetwork) InitSupplierAccountsWithSequence(
	t *testing.T,
	supplierList ...sharedtypes.Supplier,
) {
	t.Helper()

	net := memnet.GetNetwork(t)
	require.NotNil(t, net, "in-memory cosmos testutil network not initialized yet, call #Start() first")

	for _, supplier := range supplierList {
		supplierAddr, err := types.AccAddressFromBech32(supplier.GetAddress())
		require.NoError(t, err)
		memnet.InitAccount(t, supplierAddr)
	}
}

// InitAppAccountsWithSequence uses the testCLI to create on-chain accounts
// (i.e. auth module) for the addresses of the given Applications.
func (memnet *BaseInMemoryCosmosNetwork) InitAppAccountsWithSequence(
	t *testing.T,
	appList ...apptypes.Application,
) {
	t.Helper()

	for _, application := range appList {
		appAddr, err := types.AccAddressFromBech32(application.GetAddress())
		require.NoError(t, err)
		memnet.InitAccount(t, appAddr)
	}
}

// InitGatewayAccountsWithSequence uses the testCLI to create on-chain accounts
// (i.e. auth module) for the addresses of the given Gateways.
func (memnet *BaseInMemoryCosmosNetwork) InitGatewayAccountsWithSequence(
	t *testing.T,
	gatewayList ...gatewaytypes.Gateway,
) {
	t.Helper()

	for _, gateway := range gatewayList {
		gatewayAddr, err := types.AccAddressFromBech32(gateway.GetAddress())
		require.NoError(t, err)
		memnet.InitAccount(t, gatewayAddr)
	}
}