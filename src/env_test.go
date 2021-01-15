package main

import (
	"os"
	"runtime/debug"
	"strings"
	"testing"
)

func TestEnvSet(t *testing.T) {
	os.Args = []string{"cli", "config", "--current-profile", "unit_test"}
	envCli.Run()
	// Clean current profile
	os.Args = []string{"cli", "env", "clean"}
	envCli.Run()
	data := printEnv(envCli)
	if len(data) > 0 {
		t.Fail()
	}

	// Set environment values
	os.Args = []string{"cli", "env", "set", "WHERE=CLI=GITHUB", "OK=YES"}
	envCli.Run()
	data = printEnv(envCli)
	expected := `WHERE=CLI=GITHUB
OK=YES
`

	if !strings.EqualFold(string(data), expected) {
		t.Fail()
		debug.PrintStack()
	}

	// Set environment values with duplicated key
	os.Args = []string{"cli", "env", "set", "WHERE=CLI=GITHUB2"}
	envCli.Run()
	data = printEnv(envCli)
	expected = `WHERE=CLI=GITHUB
OK=YES
WHERE=CLI=GITHUB2
`
	if !strings.EqualFold(string(data), expected) {
		t.Fail()
		debug.PrintStack()
	}

	// Fix environment value, remove duplicated keys
	os.Args = []string{"cli", "env", "fix"}
	envCli.Run()
	data = printEnv(envCli)
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
	os.Args = []string{"cli", "env", "clean"}
	envCli.Run()
	data = printEnv(envCli)
	if len(data) > 0 {
		t.Fail()
	}
}
