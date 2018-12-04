package golangassert

import (
	"os"
	"testing"

	"./assert"
	"./check"
)

type Person struct {
	name string
	age  int
}

var paul = Person{
	name: "Paul",
	age:  32,
}

var peter = Person{
	name: "Peter",
	age:  21,
}

func TestEqual(t *testing.T) {
	assert.Equal(t, paul, paul, "Paul equals Paul")
	check.Equal(t, paul, paul, "Paul equals Paul")
}

func TestNotEqual(t *testing.T) {
	assert.NotEqual(t, paul, peter, "Paul does not equal Peter")
	check.NotEqual(t, paul, peter, "Paul does not equal Peter")
}

func TestNotEqualNil(t *testing.T) {
	assert.NotEqual(t, paul, nil, "Paul is not nil")
	check.NotEqual(t, paul, nil, "Paul is not nil")
}

func TestNilEqualNil(t *testing.T) {
	assert.Equal(t, nil, nil, "Nil is nil")
	check.Equal(t, nil, nil, "Nil is nil")
}

func TestIn(t *testing.T) {
	assert.In(t, "Paul does not equal Peter", "Paul", "Paul is substring of message")
	check.In(t, "Paul does not equal Peter", "Paul", "Paul is substring of message")
}

func TestError(t *testing.T) {
	// "err is not nil for unexist file"
	assert.Error(t, assert.J(os.Open("unexist.file")))
	check.Error(t, check.J(os.Open("unexist.file")))
}

func TestNilError(t *testing.T) {
	assert.NilError(t, assert.J(os.OpenFile("tempfile.txt", os.O_RDWR|os.O_CREATE, 0755)))
	check.NilError(t, check.J(os.OpenFile("tempfile.txt", os.O_RDWR|os.O_CREATE, 0755)))
	os.Remove("tempfile.txt")
}

func TestNilErr(t *testing.T) {
	assert.NilErr(t, nil)
	check.NilErr(t, nil)
}
