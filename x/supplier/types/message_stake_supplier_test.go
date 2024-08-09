package types

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/app/volatile"
	"github.com/pokt-network/poktroll/testutil/sample"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

// TODO_TECHDEBT: This test has a lot of copy-pasted code from test to test.
// It can be simplified by splitting it into smaller tests where the common
// fields don't need to be explicitly specified from test to test.
func TestMsgStakeSupplier_ValidateBasic(t *testing.T) {
	defaultServicesList := []*sharedtypes.SupplierServiceConfig{
		{
			Service: &sharedtypes.Service{
				Id: "svcId1",
			},
			Endpoints: []*sharedtypes.SupplierEndpoint{
				{
					Url:     "http://localhost:8081",
					RpcType: sharedtypes.RPCType_JSON_RPC,
					Configs: make([]*sharedtypes.ConfigOption, 0),
				},
			},
		},
	}

	ownerAddress := sample.AccAddress()
	operatorAddress := sample.AccAddress()

	tests := []struct {
		desc        string
		msg         MsgStakeSupplier
		expectedErr error
	}{
		// address related tests
		{
			desc: "valid same owner and operator address",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      ownerAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
		},
		{
			desc: "valid different owner and operator address",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
		},
		{
			desc: "valid signer is operator address",
			msg: MsgStakeSupplier{
				Signer:       operatorAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
		},
		{
			desc: "valid signer is neither the operator nor the owner",
			msg: MsgStakeSupplier{
				Signer:       sample.AccAddress(),
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
		},
		{
			desc: "invalid operator address",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      "invalid_address",
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},
		{
			desc: "invalid owner address",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: "invalid_address",
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},
		{
			desc: "invalid signer address",
			msg: MsgStakeSupplier{
				Signer:       "invalid_address",
				OwnerAddress: ownerAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(0)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},
		{
			desc: "missing owner address",
			msg: MsgStakeSupplier{
				Signer: ownerAddress,
				// OwnerAddress: ownerAddress,
				Address:  operatorAddress,
				Stake:    &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},
		{
			desc: "missing operator address",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				// Address: operatorAddress,
				Stake:    &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(0)},
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},
		{
			desc: "missing signer address",
			msg: MsgStakeSupplier{
				// Signer: ownerAddress,
				OwnerAddress: ownerAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(0)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},

		// stake related tests
		{
			desc: "valid stake",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
		},
		{
			desc: "invalid stake - missing stake",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				// Stake explicitly omitted
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		},
		{
			desc: "invalid stake - zero amount",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(0)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		},
		{
			desc: "invalid stake - negative amount",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(-100)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		},
		{
			desc: "invalid stake - invalid denom",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: "invalid", Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		},
		{
			desc: "invalid stake - missing denom",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: "", Amount: math.NewInt(100)},
				Services:     defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		},

		// service related tests
		{
			desc: "valid service configs - multiple services",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id: "svcId1",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url:     "http://localhost:8081",
								RpcType: sharedtypes.RPCType_JSON_RPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
					{
						Service: &sharedtypes.Service{
							Id: "svcId2",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url:     "http://localhost:8082",
								RpcType: sharedtypes.RPCType_GRPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
		},
		{
			desc: "invalid service configs - omitted",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				// Services explicitly omitted
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - empty",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services:     []*sharedtypes.SupplierServiceConfig{},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - invalid service ID that's too long",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id: "TooLongId1234567890",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url:     "http://localhost:8080",
								RpcType: sharedtypes.RPCType_JSON_RPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - invalid service Name that's too long",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id:   "123",
							Name: "abcdefghijklmnopqrstuvwxyzab-abcdefghijklmnopqrstuvwxyzab",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url:     "http://localhost:8080",
								RpcType: sharedtypes.RPCType_JSON_RPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - invalid service ID that contains invalid characters",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id: "12 45 !",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url:     "http://localhost:8080",
								RpcType: sharedtypes.RPCType_JSON_RPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - missing url",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id:   "svcId",
							Name: "name",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								// Url explicitly omitted
								RpcType: sharedtypes.RPCType_JSON_RPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - invalid url",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id:   "svcId",
							Name: "name",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url:     "I am not a valid URL",
								RpcType: sharedtypes.RPCType_JSON_RPC,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - missing rpc type",
			msg: MsgStakeSupplier{
				Signer:       ownerAddress,
				OwnerAddress: ownerAddress,
				Address:      operatorAddress,
				Stake:        &sdk.Coin{Denom: volatile.DenomuPOKT, Amount: math.NewInt(100)},
				Services: []*sharedtypes.SupplierServiceConfig{
					{
						Service: &sharedtypes.Service{
							Id:   "svcId",
							Name: "name",
						},
						Endpoints: []*sharedtypes.SupplierEndpoint{
							{
								Url: "http://localhost:8080",
								// RpcType explicitly omitted,
								Configs: make([]*sharedtypes.ConfigOption, 0),
							},
						},
					},
				},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		// TODO_TEST: Need to add more tests around config types
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expectedErr != nil {
				require.ErrorIs(t, err, test.expectedErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
