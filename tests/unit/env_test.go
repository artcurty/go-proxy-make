package unit

import (
	"github.com/artcurty/go-proxy-make/pkg"
	"os"
	"testing"
)

func TestGetEnvReturnsValueWhenSet(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		value        string
		defaultValue string
		want         string
	}{
		{
			name:         "Returns value when environment variable is set",
			key:          "TEST_KEY",
			value:        "test_value",
			defaultValue: "default_value",
			want:         "test_value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.key, tt.value)
			defer os.Unsetenv(tt.key)

			if got := pkg.GetEnv(tt.key, tt.defaultValue); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvReturnsDefaultWhenNotSet(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "Returns default value when environment variable is not set",
			key:          "UNSET_KEY",
			defaultValue: "default_value",
			want:         "default_value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.GetEnv(tt.key, tt.defaultValue); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvReturnsEmptyStringWhenSetToEmpty(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		value        string
		defaultValue string
		want         string
	}{
		{
			name:         "Returns empty string when environment variable is set to empty",
			key:          "EMPTY_KEY",
			value:        "",
			defaultValue: "default_value",
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.key, tt.value)
			defer os.Unsetenv(tt.key)

			if got := pkg.GetEnv(tt.key, tt.defaultValue); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
