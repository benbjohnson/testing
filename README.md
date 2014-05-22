Testing Functions for Go
========================

Below is a small collection of testing functions for Go. You don't need to import this as a dependency. Just copy these to your project as needed.

No, seriously. They're tiny functions. Just copy them.


```go
import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		v = append([]interface{}{filepath.Base(file), line}, v...)
		fmt.Printf("%s:%d: "+msg+"\n", v...)
		tb.FailNow()
	}
}

func equals(tb testing.TB, exp, act interface{}) {
	if exp != act {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: exp: %#v, got: %#v\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
```
