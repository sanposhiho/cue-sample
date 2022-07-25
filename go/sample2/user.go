package sample2

type User struct {
	// use interface{} (not exact types) to check if the type check of CUE is working.
	TEL interface{} `json:"TEL"`
	Age interface{} `json:"Age"`
}
