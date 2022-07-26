package user

import "strings"

#userinfo: X={
	TEL:          string  // TEL は string 型
	Age:          <999999 // Age は 999999 より小さい
	Is_activated: true    // Is_activated はtrue

	Age: uint // Age は uint 型.
	Age: <200 // Age は 200 より小さい.

	TEL: _ | *"000-0000-0000" // デフォルト値の定義

	_hyphen_count: strings.Count(X.TEL, "-") & 2 // TEL は二つ”-”を含む.
}
