package sample2

type User struct {
	// (CUE の型によるバリデーションが働いていることを確かめるため あえて具体的な型ではなく interface{} を使用しています。)
	TEL interface{} `json:"TEL"`
	Age interface{} `json:"Age"`
}
