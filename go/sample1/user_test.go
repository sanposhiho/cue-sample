package user

import "testing"

type user struct {
	TEL string
	Age uint
}
type testcases []struct {
	name    string
	user    user
	wantErr bool
}

var tests = testcases{
	{
		name: "OK",
		user: user{
			TEL: "000-0000-0000",
			Age: 20,
		},
		wantErr: false,
	},
	{
		name: "NG",
		user: user{
			TEL: "000-0000-0000",
			Age: 2000,
		},
		wantErr: true,
	},
}

func TestUser_Validate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &User{
				TEL: tt.user.TEL,
				Age: tt.user.Age,
			}
			if err := i.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUser_ValidateWithCUE(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &User{
				TEL: tt.user.TEL,
				Age: tt.user.Age,
			}
			if err := i.ValidateWithCUE(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
