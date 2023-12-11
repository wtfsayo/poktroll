package config_test

import (
	"net/url"
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/gogo/status"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/pkg/appgateserver/config"
	"github.com/pokt-network/poktroll/testutil/yaml"
)

func Test_ParseAppGateConfigs(t *testing.T) {
	tests := []struct {
		desc string

		inputConfig string

		expectedError  *sdkerrors.Error
		expectedConfig *config.AppGateServerConfig
	}{
		// Valid Configs
		{
			desc: "valid: AppGateServer config",

			inputConfig: `
				self_signing: true
				signing_key: app1
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				grpc_insecure: true
				pocket_node_websocket_url: tcp://127.0.0.1:36657
				`,

			expectedError: nil,
			expectedConfig: &config.AppGateServerConfig{
				SelfSigning:            true,
				SigningKey:             "app1",
				ListeningEndpoint:      &url.URL{Scheme: "http", Host: "localhost:42069"},
				QueryNodeGRPCUrl:       &url.URL{Scheme: "tcp", Host: "127.0.0.1:36658"},
				GRPCInsecure:           true,
				PocketNodeWebsocketUrl: &url.URL{Scheme: "tcp", Host: "127.0.0.1:36657", Path: "webcosket"},
			},
		},
		{
			desc: "valid: AppGateServer config with undefined self signing",

			inputConfig: `
				signing_key: app1
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				grpc_insecure: true
				pocket_node_websocket_url: tcp://127.0.0.1:36657
				`,

			expectedError: nil,
			expectedConfig: &config.AppGateServerConfig{
				SelfSigning:            false,
				SigningKey:             "app1",
				ListeningEndpoint:      &url.URL{Scheme: "http", Host: "localhost:42069"},
				QueryNodeGRPCUrl:       &url.URL{Scheme: "tcp", Host: "127.0.0.1:36658"},
				GRPCInsecure:           true,
				PocketNodeWebsocketUrl: &url.URL{Scheme: "tcp", Host: "127.0.0.1:36657", Path: "webcosket"},
			},
		},
		{
			desc: "valid: AppGateServer config with undefined grpc insecure",

			inputConfig: `
				self_signing: true
				signing_key: app1
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				pocket_node_websocket_url: tcp://127.0.0.1:36657
				`,

			expectedError: nil,
			expectedConfig: &config.AppGateServerConfig{
				SelfSigning:            true,
				SigningKey:             "app1",
				ListeningEndpoint:      &url.URL{Scheme: "http", Host: "localhost:42069"},
				QueryNodeGRPCUrl:       &url.URL{Scheme: "tcp", Host: "127.0.0.1:36658"},
				GRPCInsecure:           false,
				PocketNodeWebsocketUrl: &url.URL{Scheme: "tcp", Host: "127.0.0.1:36657", Path: "webcosket"},
			},
		},
		// Invalid Configs
		{
			desc: "invalid: empty AppGateServer config",

			inputConfig: ``,

			expectedError: config.ErrAppGateConfigUnmarshalYAML,
		},
		{
			desc: "invalid: no signing key",

			inputConfig: `
				self_signing: true
				signing_key:
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				grpc_insecure: true
				pocket_node_websocket_url: tcp://127.0.0.1:36657
				`,

			expectedError: config.ErrAppGateConfigEmptySigningKey,
		},
		{
			desc: "invalid: invalid listening endpoint",

			inputConfig: `
				self_signing: true
				signing_key: app1
				listening_endpoint: &localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				grpc_insecure: true
				pocket_node_websocket_url: tcp://127.0.0.1:36657
				`,

			expectedError: config.ErrAppGateConfigInvalidListeningEndpoint,
		},
		{
			desc: "invalid: invalid query node grpc url",

			inputConfig: `
				self_signing: true
				signing_key: app1
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: &127.0.0.1:36658
				grpc_insecure: true
				pocket_node_websocket_url: tcp://127.0.0.1:36657
				`,

			expectedError: config.ErrAppGateConfigInvalidQueryNodeGRPCUrl,
		},
		{
			desc: "invalid: invalid pocket node websocket",

			inputConfig: `
				self_signing: true
				signing_key: app1
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				grpc_insecure: true
				pocket_node_websocket_url: &127.0.0.1:36657
				`,

			expectedError: config.ErrAppGateConfigInvalidPocketNodeWebsocketUrl,
		},
		{
			desc: "invalid: missing pocket node websocket url",

			inputConfig: `
				self_signing: true
				signing_key: app1
				listening_endpoint: http://localhost:42069
				query_node_grpc_url: tcp://127.0.0.1:36658
				grpc_insecure: true
				`,

			expectedError: config.ErrAppGateConfigInvalidPocketNodeWebsocketUrl,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			normalizedConfig := yaml.NormalizeYAMLIndentation(tt.inputConfig)
			config, err := config.ParseAppGateServerConfigs([]byte(normalizedConfig))

			if tt.expectedError != nil {
				require.Error(t, err)
				require.Nil(t, config)
				stat, ok := status.FromError(tt.expectedError)
				require.True(t, ok)
				require.Contains(t, stat.Message(), tt.expectedError.Error())
				require.Nil(t, config)
				return
			}

			require.NoError(t, err)

			require.Equal(t, tt.expectedConfig.SelfSigning, config.SelfSigning)
			require.Equal(t, tt.expectedConfig.SigningKey, config.SigningKey)
			require.Equal(t, tt.expectedConfig.ListeningEndpoint.String(), config.ListeningEndpoint.String())
			require.Equal(t, tt.expectedConfig.QueryNodeGRPCUrl.String(), config.QueryNodeGRPCUrl.String())
			require.Equal(t, tt.expectedConfig.QueryNodeGRPCUrl.String(), config.QueryNodeGRPCUrl.String())
		})
	}
}
