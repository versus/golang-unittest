package assert

import "testing"

func Equal(t *testing.T, expected interface{}, actual interface{}, message string) {
	assert(t, expected, actual, message, true)
}

func NotEqual(t *testing.T, expected interface{}, actual interface{}, message string) {
	assert(t, expected, actual, message, false)
}

func assert(t *testing.T, expected interface{}, actual interface{}, message string, expectation bool) {
	if (expected == actual) != expectation {
		t.Errorf("%s failed: Expected %v, but got %v", message, expected, actual)
	}
}
