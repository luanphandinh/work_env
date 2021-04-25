package main

import (
	"os"
	"testing"

	"github.com/luanphandinh/env/src/assert"
)

func TestConfigSetCurrentProfile(t *testing.T) {
	os.Args = []string{"cli", "set", "config", "--current-profile", "unit_test"}
	envCli.Run()
	assert.StringNotEquals(t, envCli.GetConfig("current_profile").(string), "default")
	assert.StringEquals(t, envCli.GetConfig("current_profile").(string), "unit_test")

	os.Args = []string{"cli", "set", "config", "--current-profile", "default"}
	envCli.Run()
	assert.StringEquals(t, envCli.GetConfig("current_profile").(string), "default")
	assert.StringNotEquals(t, envCli.GetConfig("current_profile").(string), "unit_test")

	// Should run with optional --p arguments
	os.Args = []string{"cli", "--p", "unit_test_2", "describe", "config"}
	envCli.Run()
	assert.StringEquals(t, envCli.GetConfig("current_profile").(string), "unit_test_2")
}
