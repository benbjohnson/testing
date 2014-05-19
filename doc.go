/*

You have to be kidding me. It's literally two functions. Why are you reading
the documentation?

	func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
		if !condition {
			tb.Fatalf(msg, v...)
		}
	}

	func equals(tb testing.TB, exp, act interface{}) {
		assert(tb, exp == act, "exp: %#v, got: %#v", exp, act)
	}

*/
package testing
