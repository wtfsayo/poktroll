package supplier

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/proto/types/shared"
	"github.com/pokt-network/poktroll/testutil/sample"
)

// TODO_TECHDEBT: This test has a lot of copy-pasted code from test to test.
// It can be simplified by splitting it into smaller tests where the common
// fields don't need to be explicitly specified from test to test.
func TestMsgStakeSupplier_ValidateBasic(t *testing.T) {
	defaultServicesList := []*shared.SupplierServiceConfig{
		{
			Service: &shared.Service{
				Id: "svcId1",
			},
			Endpoints: []*shared.SupplierEndpoint{
				{
					Url:     "http://localhost:8081",
					RpcType: shared.RPCType_JSON_RPC,
					Configs: make([]*shared.ConfigOption, 0),
				},
			},
		}}

	tests := []struct {
		desc        string
		msg         MsgStakeSupplier
		expectedErr error
	}{
		// address related tests
		{
			desc: "invalid address - nil stake",
			msg: MsgStakeSupplier{
				Address: "invalid_address",
				// Stake explicitly omitted
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidAddress,
		},

		// stake related tests
		{
			desc: "valid address - nil stake",
			msg: MsgStakeSupplier{
				Address: sample.AccAddress(),
				// Stake explicitly omitted
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		}, {
			desc: "valid address - valid stake",
			msg: MsgStakeSupplier{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: defaultServicesList,
			},
		}, {
			desc: "valid address - zero stake",
			msg: MsgStakeSupplier{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "upokt", Amount: math.NewInt(0)},
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		}, {
			desc: "valid address - negative stake",
			msg: MsgStakeSupplier{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "upokt", Amount: math.NewInt(-100)},
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		}, {
			desc: "valid address - invalid stake denom",
			msg: MsgStakeSupplier{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "invalid", Amount: math.NewInt(100)},
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		}, {
			desc: "valid address - invalid stake missing denom",
			msg: MsgStakeSupplier{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "", Amount: math.NewInt(100)},
				Services: defaultServicesList,
			},
			expectedErr: ErrSupplierInvalidStake,
		},

		// service related tests
		{
			desc: "valid service configs - multiple services",
			msg: MsgStakeSupplier{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id: "svcId1",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url:     "http://localhost:8081",
								RpcType: shared.RPCType_JSON_RPC,
								Configs: make([]*shared.ConfigOption, 0),
							},
						},
					},
					{
						Service: &shared.Service{
							Id: "svcId2",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url:     "http://localhost:8082",
								RpcType: shared.RPCType_GRPC,
								Configs: make([]*shared.ConfigOption, 0),
							},
						},
					},
				},
			},
		},
		{
			desc: "invalid service configs - omitted",
			msg: MsgStakeSupplier{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				// Services explicitly omitted
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - empty",
			msg: MsgStakeSupplier{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{},
			},
			expectedErr: ErrSupplierInvalidServiceConfig,
		},
		{
			desc: "invalid service configs - invalid service ID that's too long",
			msg: MsgStakeSupplier{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id: "TooLongId1234567890",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url:     "http://localhost:8080",
								RpcType: shared.RPCType_JSON_RPC,
								Configs: make([]*shared.ConfigOption, 0),
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
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id:   "123",
							Name: "abcdefghijklmnopqrstuvwxyzab-abcdefghijklmnopqrstuvwxyzab",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url:     "http://localhost:8080",
								RpcType: shared.RPCType_JSON_RPC,
								Configs: make([]*shared.ConfigOption, 0),
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
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id: "12 45 !",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url:     "http://localhost:8080",
								RpcType: shared.RPCType_JSON_RPC,
								Configs: make([]*shared.ConfigOption, 0),
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
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id:   "svcId",
							Name: "name",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								// Url explicitly omitted
								RpcType: shared.RPCType_JSON_RPC,
								Configs: make([]*shared.ConfigOption, 0),
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
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id:   "svcId",
							Name: "name",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url:     "I am not a valid URL",
								RpcType: shared.RPCType_JSON_RPC,
								Configs: make([]*shared.ConfigOption, 0),
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
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: math.NewInt(100)},
				Services: []*shared.SupplierServiceConfig{
					{
						Service: &shared.Service{
							Id:   "svcId",
							Name: "name",
						},
						Endpoints: []*shared.SupplierEndpoint{
							{
								Url: "http://localhost:8080",
								// RpcType explicitly omitted,
								Configs: make([]*shared.ConfigOption, 0),
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
