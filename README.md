# Testa Rossa

<p align="center"><img src="testarossa.png" width="256"></p><p></p>

`TestaRossa` is a simple utility on top of Go's standard `testing` library that prints the `full/path/to/the/source/file/of/the/unit_test.go:line` of a failed assertion, which in VSCode is a clickable link. Unlike `t.Log`, output is printed using `fmt.Printf` during the execution of the test rather than at its conclusion.

```go
func TestMe(t *testing.T) {
	testarossa.Equal(t, 1, 0, "You are not the one")
	err := errors.New("This is bad")
	testarossa.NoError(t, err)

	droids := 1234
	testarossa.FailIf(t, droids != 0, "These are not the droids you are looking for")
	err = errors.New("This is really bad")
	testarossa.FatalIfError(t, err)
}
```

outputs

```
--- FAIL: TestMe
    /my_projects/go/github.com/microbus-io/testarossa/fail_test.go:27
    Expected 1, actual 0
    You are not the one
--- FAIL: TestMe
    /my_projects/go/github.com/microbus-io/testarossa/fail_test.go:29
    Expected no error
    This is bad
--- FAIL: TestMe
    /my_projects/go/github.com/microbus-io/testarossa/fail_test.go:32
    These are not the droids you are looking for
--- FAIL: TestMe
    /my_projects/go/github.com/microbus-io/testarossa/fail_test.go:34
    This is really bad
--- FAIL: TestMe (0.00s)
FAIL
FAIL	github.com/microbus-io/testarossa	0.193s
FAIL
```

`TestaRossa` is licensed by Microbus LLC under the [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0).
