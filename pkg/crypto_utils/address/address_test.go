package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAddress(t *testing.T) {
	tests := []struct {
		address  string
		expected bool
	}{
		{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e", true},  // valid address
		{"742d35Cc6634C0532925a3b844Bc454e4438f44e", false},   // missing "0x" prefix
		{"0x742d35Cc6634C0532925a3b844Bc454e4438f44", false},  // too short
		{"0xG42d35Cc6634C0532925a3b844Bc454e4438f44e", false}, // contains invalid character "G"
		{"", false}, // empty string
	}

	for _, tt := range tests {
		result := CheckAddress(tt.address)
		assert.Equal(t, tt.expected, result)
	}
}
