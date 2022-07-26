package user

import (
	"errors"

	"cuelang.org/go/encoding/gocode/gocodec"
)

type User struct {
	TEL string
	Age uint `cue:"<200"`
}

var ErrInvalidUser = errors.New("invalid user")

// 通常のGoにおけるバリデーションの実装例
func (i *User) Validate() error {
	if i.Age >= 200 {
		return ErrInvalidUser
	}

	return nil
}

// cueタグに書かれた制約をもとにしたバリデーションの実装例
func (i *User) ValidateWithCUE() error {
	return gocodec.Validate(i)
}
