package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams_ValidateNumBlocksPerSession(t *testing.T) {
	tests := []struct {
		desc                string
		numBlocksPerSession any
		err                 error
	}{
		{
			desc:                "invalid type",
			numBlocksPerSession: "invalid",
			err:                 ErrSessionParamInvalid.Wrapf("invalid parameter type: %T", "invalid"),
		},
		{
			desc:                "zero NumBlocksPerSession",
			numBlocksPerSession: uint64(0),
			err:                 ErrSessionParamInvalid.Wrapf("invalid NumBlocksPerSession: (%v)", uint64(0)),
		},
		{
			desc:                "valid NumBlocksPerSession",
			numBlocksPerSession: uint64(4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := ValidateNumBlocksPerSession(tt.numBlocksPerSession)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
