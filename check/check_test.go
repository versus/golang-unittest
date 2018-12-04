package check

import (
	"os"
	"testing"
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
	Equal(t, paul, paul, "Paul equals Paul")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, paul, peter, "Paul does not equal Peter")
}

func TestNotEqualNil(t *testing.T) {
	NotEqual(t, paul, nil, "Paul is not nil")
}

func TestNilEqualNil(t *testing.T) {
	Equal(t, nil, nil, "Nil is nil")
}

func TestIn(t *testing.T) {
	In(t, "Paul does not equal Peter", "Paul", "Paul is substring of message")
}

func TestError(t *testing.T) {
	// "err is not nil for unexist file"
	Error(t, J(os.Open("unexist.file")))
}

func TestNilError(t *testing.T) {
	NilError(t, J(os.OpenFile("tempfile.txt", os.O_RDWR|os.O_CREATE, 0755)))
	os.Remove("tempfile.txt")
}

func TestNilErr(t *testing.T) {
	NilErr(t, nil)
}
