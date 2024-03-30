package config

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConf(t *testing.T) {
	// rand int
	randInt := rand.Intn(10000000)

	// Set the environment variables
	os.Setenv("MAX_CONN", fmt.Sprintf("%d", randInt))

	// Simulate the function call
	NewConf()

	assert.Equal(t, "314159265358979323846264338327950288419716939937510582097494459", Conf.Key)
	assert.Equal(t, randInt, Conf.MaxConn)
}

func TestConfigValidate(t *testing.T) {
	cfg := &Config{}

	// Check for the invalid configuration
	assert.False(t, cfg.Validate())

	// Set the required fields
	cfg.ID = "1"
	cfg.Moniker = "tester"
	cfg.Key = "secretkey"
	cfg.IP = "127.0.0.1"
	cfg.PartyPassword = "partypass"

	// Now the configuration should be valid
	assert.True(t, cfg.Validate())
}
