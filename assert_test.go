package assert

import (
	"testing"
)

type Person struct {
	name string
	age  int
}

var paul = Person{
	name: "Paul",
	age:  32}

var peter = Person{
	name: "Peter",
	age:  21}

func TestEqual(t *testing.T) {
	Equal(t, paul, paul, "Paul equals paul")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, paul, peter, "Paul does not equal peter")
}
