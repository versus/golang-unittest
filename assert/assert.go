package assert

import (
	"runtime"
	"strings"
	"testing"
)

func assert(t *testing.T, actual interface{}, expected interface{}, message string, expectation bool) {
	if (expected == actual) != expectation {
		if message == "" {
			message = "Test"
		}
		buf := make([]byte, 1024)
		runtime.Stack(buf, false)
		tracepieces := strings.Split(string(buf), "\n")

		// 6-th points to assert-line in test class
		if expectation {
			t.Fatalf("from %s: %s failed: Expected |%v|, but got |%v|", tracepieces[6], message, expected, actual)
		} else {
			t.Fatalf("from %s: %s failed: Expected not |%v|, but got |%v|", tracepieces[6], message, expected, actual)
		}
	}
}

// Equal checks two parameters have the same type and value
func Equal(t *testing.T, actual interface{}, expected interface{}, message string) {
	assert(t, actual, expected, message, true)
}

// In checks the second string parameter is in the first string parameter
func In(t *testing.T, actual string, substring string, message string) {
	if strings.Index(actual, substring) < 0 {
		if message == "" {
			message = "Test"
		}
		t.Fatalf("%s failed: Expected substring |%v| in |%v|", message, substring, actual)
	}
}

// NotEqual checks two parameters have the different type or value
func NotEqual(t *testing.T, actual interface{}, expected interface{}, message string) {
	assert(t, actual, expected, message, false)
}

func Error(t *testing.T, params []interface{}) {
	assert(t, params[1], nil, "assert.Error", false)
}

func NilError(t *testing.T, params []interface{}) {
	assert(t, params[1], nil, "assert.NilError", true)
}

func NilErr(t *testing.T, err error) {
	assert(t, err, nil, "assert.NilErr", true)
}

func J(a ...interface{}) []interface{} {
	return a
}
