package main

import (
	"env/src/assert"
	"os"
	"testing"
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
}
