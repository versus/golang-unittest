package check

import (
	"runtime"
	"strings"
	"testing"
)

func check(t *testing.T, actual interface{}, expected interface{}, message string, expectation bool) {
	if (expected == actual) != expectation {
		if message == "" {
			message = "Test"
		}
		buf := make([]byte, 1024)
		runtime.Stack(buf, false)
		tracepieces := strings.Split(string(buf), "\n")

		// 6-th points to check-line in test class
		t.Errorf("from %s: %s failed: Expected |%v|, but got |%v|", tracepieces[6], message, expected, actual)
	}
}

func Equal(t *testing.T, actual interface{}, expected interface{}, message string) {
	check(t, actual, expected, message, true)
}

func In(t *testing.T, actual string, substring string, message string) {
	if strings.Index(actual, substring) < 0 {
		if message == "" {
			message = "Test"
		}
		t.Errorf("%s failed: Expected substring |%v| in |%v|", message, substring, actual)
	}
}

func NotEqual(t *testing.T, actual interface{}, expected interface{}, message string) {
	check(t, actual, expected, message, false)
}

func Error(t *testing.T, params []interface{}) {
	check(t, params[1], nil, "check.Error", false)
}

func NilError(t *testing.T, params []interface{}) {
	check(t, params[1], nil, "check.NilError", true)
}

func NilErr(t *testing.T, err error) {
	check(t, err, nil, "check.NilErr", true)
}

func J(a ...interface{}) []interface{} {
	return a
}
