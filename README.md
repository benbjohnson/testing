Testing Functions for Go
========================

Below is a small collection of testing functions for Go. You don't need to import this as a dependency. Just copy these to your project as needed.

No, seriously. They're tiny functions. Just copy them.


```go
import (
	"fmt"
	"path/filepath"
	"runtime"
	"reflect"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if got is not equal to want.
func equals(tb testing.TB, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\n\n\twant: %#v\033[39m\n\n", filepath.Base(file), line, got, want)
		tb.FailNow()
	}
}

```
