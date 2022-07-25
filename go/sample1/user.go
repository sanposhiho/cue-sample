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

// Gophers often define the validation like this.
func (i *User) Validate() error {
	if i.Age >= 200 {
		return ErrInvalidUser
	}

	return nil
}

// But, we can validate it with cue as well!
func (i *User) ValidateWithCUE() error {
	return gocodec.Validate(i)
}
