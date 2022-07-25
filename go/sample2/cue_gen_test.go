package sample2

import "testing"

func TestUser_Validate(t *testing.T) {
	type user struct {
		TEL interface{}
		Age interface{}
	}
	tests := []struct {
		name    string
		user    user
		wantErr bool
	}{
		{
			name: "OK",
			user: user{
				TEL: "999-9999-9999",
				Age: 10,
			},
		},
		{
			name: "NG: no '-'",
			user: user{
				TEL: "invalid",
				Age: 10000,
			},
			wantErr: true,
		},
		{
			name: "NG: invalid count",
			user: user{
				TEL: "000-000-000",
				Age: 10000,
			},
			wantErr: true,
		},
		{
			name: "NG: too many '-'",
			user: user{
				TEL: "0-0-0-0-0",
				Age: 10000,
			},
			wantErr: true,
		},
		{
			name: "NG: type is invalid",
			user: user{
				TEL: 10000, // should be string.
				Age: "age", // should be uint type.
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &User{
				TEL: tt.user.TEL,
				Age: tt.user.Age,
			}
			if err := x.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
