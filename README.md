# golang asserts and checks

[![Build Status](https://travis-ci.org/vizor-games/golang-unittest.svg)](https://travis-ci.org/vizor-games/golang-unittest)

Two set of assert and check functions for hard and soft unit test assertions in Go. The golang creators reject to create own testing framework, see [Go FAQ about testing framework](https://golang.org/doc/faq#testing_framework).

Insipered by:
* [golang-assert](https://github.com/stfsy/golang-assert)
* [python unittest library](https://docs.python.org/3.6/library/unittest.html)

Improvements:
* check functions
* frendlier error messages, including correct code line from tests
* In function for strings
* Error, NilError functions for functions that returns more than one parameter
* NilErr - short way to check no errors

The check functions are soft assertions. Failing tests will only be reported to the console via *t.Errorf()*.
The assert functions are hard assertions. Failing tests will be reported to the console via *t.Failf()* and skip the current test execution.


## Check equal

```go
import "https://github.com/vizor-games/golang-unittest/check"

type Person struct {
	name string
	age  int
}

var paul = Person{
	name: "Paul",
	age:  32,
}

paul2 := paul

func TestEqual(t *testing.T) {
	check.Equal(t, paul, paul2, "Paul equals Paul")
}
```

## Assert not equal

```go
var peter = Person{
	name: "Peter",
	age:  21,
}

func TestNotEqual(t *testing.T) {
	assert.NotEqual(t, paul, peter, "Paul does not equal Peter")
```

## Functions that returns additional error value

Use assert.J or check.J to convert args... to []interface{}:

```go
func TestError(t *testing.T) {
	// "err is not nil for unexist file"
	assert.Error(t, assert.J(os.Open("unexist.file")))
}
```

## Installation

New way:
```bash
vgo get github.com/vizor-games/golang-unittest
```

Old way:
```bash
go get github.com/vizor-games/golang-unittest
```

## Run tests

```bash
go test -v  # -v - verbose
go test ./assert
go test ./check
```

## License

This project is distributed under the MIT license.
