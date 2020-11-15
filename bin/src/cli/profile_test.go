package main

import (
	"runtime/debug"
	"strings"
	"testing"
)

func TestProfileSet(t *testing.T) {
	cli := &CLI{}
	cli.LoadConfig()
	cli.Config.CurrentProfile = "unit_test"

	// Clean current profile
	CleanProfileEnv(cli)
	data := GetProfileEnv(cli)
	if len(data) > 0 {
		t.Fail()
	}

	// Set environment values
	cli.args = []string{"WHERE=CLI=GITHUB", "OK=YES"}
	SetProfileEnv(cli)
	data = GetProfileEnv(cli)
	expected := `WHERE=CLI=GITHUB
OK=YES
`

	if !strings.EqualFold(string(data), expected) {
		t.Fail()
		debug.PrintStack()
	}

	// Set environment values with duplicated key
	cli.args = []string{"WHERE=CLI=GITHUB2"}
	SetProfileEnv(cli)
	data = GetProfileEnv(cli)
	expected = `WHERE=CLI=GITHUB
OK=YES
WHERE=CLI=GITHUB2
`
	if !strings.EqualFold(string(data), expected) {
		t.Fail()
		debug.PrintStack()
	}


	// Fix environment value, remove duplicated keys
	FixProfileEnv(cli)
	data = GetProfileEnv(cli)
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
	CleanProfileEnv(cli)
	data = GetProfileEnv(cli)
	if len(data) > 0 {
		t.Fail()
	}
}
