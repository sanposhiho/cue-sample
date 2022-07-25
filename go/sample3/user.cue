package sample3

import "strings"

#User: X={
	_hyphen_count: strings.Count(X.TEL, "-") & 2 // should contain two "-"
	_length: len(X.TEL) & (12 | 13) // the length of TEL in japan is 10 or 11

	TEL: string

	Age: <999
	Age: uint
}
