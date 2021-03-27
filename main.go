package main

import (
	"os"
	"protoc-gen-rest/parse"
	"protoc-gen-rest/plugin/ts"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	//for debug
	// write data input to file for debug

	// data, _ := ioutil.ReadAll(os.Stdin)
	// ioutil.WriteFile(".scalars.txt", data, 0644)

	// pgs.Init(
	// 	pgs.ProtocInput(os.Stdin),
	// ).RegisterModule(
	// 	ts.TsGen(),
	// ).RegisterPostProcessor().Render()

	tsParser := parse.NewTsParser()
	f, _ := os.Open(".scalars.txt")
	pgs.Init(
		pgs.ProtocInput(f),
	).RegisterModule(
		ts.TsGen(tsParser),
	).RegisterPostProcessor().Render()
}
