package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     Config{Port: "8080", LoggerLevel: "INFO"},
			wantErr: false,
		},
		{
			name:    "missing port",
			cfg:     Config{Port: "", LoggerLevel: "INFO"},
			wantErr: true,
		},
		{
			name:    "invalid port (letters)",
			cfg:     Config{Port: "abc", LoggerLevel: "INFO"},
			wantErr: true,
		},
		{
			name:    "invalid logger level",
			cfg:     Config{Port: "8080", LoggerLevel: "VERBOSE"},
			wantErr: true,
		},
		{
			name:    "empty config",
			cfg:     Config{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
