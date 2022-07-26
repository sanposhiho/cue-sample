./sample1/user.go から ./sample1/user_go_gen.cue の生成

```shell
cue get go github.com/sanposhiho/cue-sample/go/sample1/ --local
```

テストの実行

```shell
go test ./sample1/...
```
