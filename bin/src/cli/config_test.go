package main

import (
	"testing"
)

func TestConfigSetCurrentProfile(t *testing.T) {
	cli := &CLI{}
	cli.Init()

	cli.Config.CurrentProfile = "unit_test"
	cli.SaveConfig()
	cli.LoadConfig()
	if cli.Config.CurrentProfile != "unit_test" {
		t.Fail()
	}

	cli.Config.CurrentProfile = "unit_test_again"
	cli.LoadConfig()
	if cli.Config.CurrentProfile != "unit_test" {
		t.Fail()
	}

	cli.Config.CurrentProfile = "unit_test_again"
	cli.SaveConfig()
	cli.LoadConfig()
	if cli.Config.CurrentProfile != "unit_test_again" {
		t.Fail()
	}
}
