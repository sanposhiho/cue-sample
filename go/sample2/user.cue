package sample2

import "strings"

User: X={
	_hyphen_count: strings.Count(X.TEL, "-") & 2 // TEL は二つ"-"を含む
	_length: len(X.TEL) & (12 | 13) // TEL の長さは 10 or 11

	TEL: string

	Age: <999
	Age: uint
}

