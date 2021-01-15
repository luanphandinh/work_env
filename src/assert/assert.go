package assert

import (
	"runtime/debug"
	"testing"
)

func StringEquals(t *testing.T, s1, s2 string) {
	if s1 != s2 {
		t.Fail()
		debug.PrintStack()
	}
}

func StringNotEquals(t *testing.T, s1, s2 string) {
	if s1 == s2 {
		t.Fail()
		debug.PrintStack()
	}
}
