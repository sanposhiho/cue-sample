package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"cuelang.org/go/cue/load"

	"cuelang.org/go/cue"

	"cuelang.org/go/encoding/gocode"
)

const (
	cueFilePath = "./sample2/user.cue"

	module      = "github.com/sanposhiho/cue-sample/go/sample2"
	dir         = "./sample2/"
	destination = "./sample2/cue_gen.go"
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

	b, err = gocode.Generate(dir, inst, nil)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(destination, b, 0644)
	if err != nil {
		panic(err)
	}
}
