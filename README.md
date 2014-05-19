Testing Functions for Go
========================

Below is a small collection of testing functions for Go. You don't need to import this as a dependency. Just copy these to your project as needed.

No, seriously. They're tiny functions. Just copy them.


```go
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		tb.Fatalf(msg, v...)
	}
}

func equals(tb testing.TB, exp, act interface{}) {
	assert(tb, exp == act, "exp: %#v, got: %#v", exp, act)
}
```
