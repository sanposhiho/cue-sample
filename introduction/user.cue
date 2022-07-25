package user

import "strings"

#userinfo: X={
	TEL:          string  // TEL should be string type.
	Age:          <999999 // Age should be <999999.
	Is_activated: true    // Is_activated should be true.

	_hyphen_count: strings.Count(X.TEL, "-") & 2 // TEL should contain two "-".
	_length:       len(X.TEL) & (12 | 13)        // the length of TEL should be 10 or 11.

	Age: uint // Age should be uint type.
	Age: <200 // Age should be <200.
}
