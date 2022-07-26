./sample2/user.cue から ./sample2/cue_gen.go の生成

```shell
# ./basic/sample2/generator/go/generate.go で、user.cue から cue_gen.go を生成
go run ./sample2/generator/go/generate.go
# テストの実行
go test ./sample2/...
```
