generate ./sample1/user_go_gen.cue from ./sample1/user.go

```shell
cue get go github.com/sanposhiho/cue-sample1/go/sample1/ --local
# run test
go test ./sample1/...
```
