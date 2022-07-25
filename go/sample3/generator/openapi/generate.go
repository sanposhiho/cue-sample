package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"cuelang.org/go/encoding/openapi"

	"cuelang.org/go/cue/load"

	"cuelang.org/go/cue"
)

const (
	cueFilePath = "./sample3/user.cue"

	module      = "github.com/sanposhiho/cue-sample/go/sample3"
	dir         = "./sample3/"
	destination = "./sample3/openapi.json"
)

func main() {
	b, err := os.ReadFile(cueFilePath)
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir := filepath.Join(cwd, dir)
	inst := cue.Build(load.Instances([]string{module}, &load.Config{
		Dir:        dir,
		ModuleRoot: dir,
		Module:     module,
	}))[0]
	if err := inst.Err; err != nil {
		panic(err)
	}

	b, err = openapi.Gen(inst, nil)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "   ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(destination, out.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
