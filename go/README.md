generate ./sample1/user_go_gen.cue from ./sample1/user.go

```shell
cue get go github.com/sanposhiho/cue-sample1/go/sample1/ --local
# run test
go test ./sample1/...
```

generate ./sample2/cue_gen.go from ./sample2/user.cue

```shell
# ./basic/sample2/generator/go/generate.go has the code to generate go code from CUE.
go run ./sample2/generator/go/generate.go
# run test
go test ./sample2/...
```

generate ./sample3/openapi.json from ./sample3/user.cue

```shell
# ./basic/sample3/generator/openapi/generate.go has the code to generate openapi from CUE.
go run ./sample3/generator/openapi/generate.go
```

