generate ./sample2/cue_gen.go from ./sample2/user.cue

```shell
# ./basic/sample2/generator/go/generate.go has the code to generate go code from CUE.
go run ./sample2/generator/go/generate.go
# run test
go test ./sample2/...
```
