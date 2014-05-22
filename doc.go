/*

You have to be kidding me.

It's literally two functions.

Why are you reading the documentation?

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

*/
package testing
