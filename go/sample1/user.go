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

func (i *User) Validate() error {
	if i.Age >= 200 {
		return ErrInvalidUser
	}

	return nil
}

// But, we can validate with cue!
func (i *User) ValidateWithCUE() error {
	return gocodec.Validate(i)
}
