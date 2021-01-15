package main

import (
	"runtime/debug"
	"strings"
	"testing"
)

func TestEnvSet(t *testing.T) {
	cli := &CLI{}
	cli.LoadConfig()
	cli.Config.CurrentProfile = "unit_test"

	// Clean current profile
	CleanEnv(cli)
	data := PrintEnv(cli)
	if len(data) > 0 {
		t.Fail()
	}

	// Set environment values
	cli.args = []string{"WHERE=CLI=GITHUB", "OK=YES"}
	SetEnv(cli)
	data = PrintEnv(cli)
	expected := `WHERE=CLI=GITHUB
OK=YES
`

	if !strings.EqualFold(string(data), expected) {
		t.Fail()
		debug.PrintStack()
	}

	// Set environment values with duplicated key
	cli.args = []string{"WHERE=CLI=GITHUB2"}
	SetEnv(cli)
	data = PrintEnv(cli)
	expected = `WHERE=CLI=GITHUB
OK=YES
WHERE=CLI=GITHUB2
`
	if !strings.EqualFold(string(data), expected) {
		t.Fail()
		debug.PrintStack()
	}

	// Fix environment value, remove duplicated keys
	FixEnv(cli)
	data = PrintEnv(cli)
	expected1 := `WHERE=CLI=GITHUB2
OK=YES
`
	expected2 := `OK=YES
WHERE=CLI=GITHUB2
`
	if !strings.EqualFold(string(data), expected1) && !strings.EqualFold(string(data), expected2) {
		t.Fail()
		debug.PrintStack()
	}

	// Clean profile again
	CleanEnv(cli)
	data = PrintEnv(cli)
	if len(data) > 0 {
		t.Fail()
	}
}
