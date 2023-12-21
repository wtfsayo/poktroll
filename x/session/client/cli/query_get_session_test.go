package cli_test

import (
	"context"
	"fmt"
	"testing"

	sdkerrors "cosmossdk.io/errors"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/gogo/status"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/testutil/network"
	"github.com/pokt-network/poktroll/testutil/network/sessionnet"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	"github.com/pokt-network/poktroll/x/session/client/cli"
	sessiontypes "github.com/pokt-network/poktroll/x/session/types"
	suppliertypes "github.com/pokt-network/poktroll/x/supplier/types"
)

func TestCLI_GetSession(t *testing.T) {
	ctx := context.Background()
	// Prepare the network
	appSupplierPairingRatio := 2
	memnet := sessionnet.NewInMemoryNetworkWithSessions(
		t, &network.InMemoryNetworkConfig{
			NumSuppliers:            2,
			AppSupplierPairingRatio: appSupplierPairingRatio,
		},
	)
	memnet.Start(ctx, t)

	net := memnet.GetNetwork(t)

	// TODO_DISCUSS_IN_THIS_COMMIT: do we still need this?
	_, err := net.WaitForHeight(10) // Wait for a sufficiently high block height to ensure the staking transactions have been processed
	require.NoError(t, err)

	appGenesisState := network.GetGenesisState[*apptypes.GenesisState](t, apptypes.ModuleName, memnet)
	applications := appGenesisState.ApplicationList

	supplierGenesisState := network.GetGenesisState[*suppliertypes.GenesisState](t, suppliertypes.ModuleName, memnet)
	suppliers := supplierGenesisState.SupplierList

	// Sanity check the application configs are what we expect them to be
	appSvc0 := applications[0]
	appSvc1 := applications[2]

	require.Len(t, appSvc0.ServiceConfigs, 2)
	require.Len(t, appSvc1.ServiceConfigs, 2)

	require.Equal(t, appSvc0.ServiceConfigs[0].Service.Id, "svc0")   // svc0 has a supplier
	require.Equal(t, appSvc0.ServiceConfigs[1].Service.Id, "nosvc0") // nosvc0 doesn't have a supplier
	require.Equal(t, appSvc1.ServiceConfigs[0].Service.Id, "svc1")   // svc1 has a supplier
	require.Equal(t, appSvc1.ServiceConfigs[1].Service.Id, "nosvc1") // nosvc1 doesn't have a supplier

	// Sanity check the supplier configs are what we expect them to be
	supplierSvc0 := suppliers[0] // supplier for svc0
	supplierSvc1 := suppliers[1] // supplier for svc1

	require.Len(t, supplierSvc0.Services, 1)
	require.Len(t, supplierSvc1.Services, 1)

	require.Equal(t, supplierSvc0.Services[0].Service.Id, "svc0")
	require.Equal(t, supplierSvc1.Services[0].Service.Id, "svc1")

	// Prepare the test cases
	tests := []struct {
		desc string

		appAddress  string
		serviceId   string
		blockHeight int64

		expectedErr          *sdkerrors.Error
		expectedNumSuppliers int
	}{
		// Valid requests
		{
			desc: "valid - block height specified and is zero",

			appAddress:  appSvc0.Address,
			serviceId:   "svc0",
			blockHeight: 0,

			expectedErr:          nil,
			expectedNumSuppliers: 1,
		},
		{
			desc: "valid - block height specified and is greater than zero",

			appAddress:  appSvc1.Address,
			serviceId:   "svc1",
			blockHeight: 10,

			expectedErr:          nil,
			expectedNumSuppliers: 1,
		},
		{
			desc: "valid - block height unspecified and defaults to 0",

			appAddress: appSvc0.Address,
			serviceId:  "svc0",
			// blockHeight: intentionally omitted,

			expectedErr:          nil,
			expectedNumSuppliers: 1,
		},

		// Invalid requests - incompatible state
		{
			desc: "invalid - app not staked for service",

			appAddress:  appSvc0.Address,
			serviceId:   "svc9001", // appSvc0 is only staked for svc0 (has supplier) and svc00 (doesn't have supplier) and is not staked for service over 9000
			blockHeight: 0,

			expectedErr: sessiontypes.ErrSessionAppNotStakedForService,
		},
		{
			desc: "invalid - no suppliers staked for service",

			appAddress:  appSvc0.Address, // dynamically getting address from applications
			serviceId:   "svc00",         // appSvc0 is only staked for svc0 (has supplier) and svc00 (doesn't have supplier)
			blockHeight: 0,

			expectedErr: sessiontypes.ErrSessionSuppliersNotFound,
		},
		{
			desc: "invalid - block height is in the future",

			appAddress:  appSvc0.Address, // dynamically getting address from applications
			serviceId:   "svc0",
			blockHeight: 9001, // block height over 9000 is greater than the context height of 10

			expectedErr: sessiontypes.ErrSessionInvalidBlockHeight,
		},

		// Invalid requests - bad app address input
		{
			desc: "invalid - invalid appAddress",

			appAddress:  "invalidAddress", // providing a deliberately invalid address
			serviceId:   "svc0",
			blockHeight: 0,

			expectedErr: sessiontypes.ErrSessionInvalidAppAddress,
		},
		{
			desc: "invalid - missing appAddress",
			// appAddress: intentionally omitted
			serviceId:   "svc0",
			blockHeight: 0,

			expectedErr: sessiontypes.ErrSessionInvalidAppAddress,
		},

		// Invalid requests - bad serviceID input
		{
			desc:        "invalid - invalid service ID",
			appAddress:  appSvc0.Address, // dynamically getting address from applications
			serviceId:   "invalidServiceId",
			blockHeight: 0,

			expectedErr: sessiontypes.ErrSessionInvalidService,
		},
		{
			desc:       "invalid - missing service ID",
			appAddress: appSvc0.Address, // dynamically getting address from applications
			// serviceId:   intentionally omitted
			blockHeight: 0,

			expectedErr: sessiontypes.ErrSessionInvalidService,
		},
	}

	// We want to use the `--output=json` flag for all tests so it's easy to unmarshal below
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			// Prepare the arguments for the CLI command
			args := []string{
				tt.appAddress,
				tt.serviceId,
				fmt.Sprintf("%d", tt.blockHeight),
			}
			args = append(args, common...)

			// Execute the command
			getSessionOut, err := clitestutil.ExecTestCLICmd(memnet.GetClientCtx(t), cli.CmdGetSession(), args)
			if tt.expectedErr != nil {
				stat, ok := status.FromError(tt.expectedErr)
				require.True(t, ok)
				require.Contains(t, stat.Message(), tt.expectedErr.Error())
				return
			}
			require.NoError(t, err)

			var getSessionRes sessiontypes.QueryGetSessionResponse
			err = net.Config.Codec.UnmarshalJSON(getSessionOut.Bytes(), &getSessionRes)
			require.NoError(t, err)
			require.NotNil(t, getSessionRes)

			session := getSessionRes.Session
			require.NotNil(t, session)

			// Verify some data about the session
			require.Equal(t, tt.appAddress, session.Application.Address)
			require.Equal(t, tt.serviceId, session.Header.Service.Id)
			require.Len(t, session.Suppliers, tt.expectedNumSuppliers)
		})
	}
}
