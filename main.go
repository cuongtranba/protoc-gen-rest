package main

import (
	"os"
	"protoc-gen-rest/parse"
	"protoc-gen-rest/plugin/gopl"
	"protoc-gen-rest/plugin/ts"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	//for debug
	// write data input to file for debug
	// data, _ := ioutil.ReadAll(os.Stdin)
	// ioutil.WriteFile(".scalars.txt", data, 0644)

	pgs.Init(
		pgs.ProtocInput(os.Stdin),
	).RegisterModule(
		gopl.GoGen(),
		ts.TsGen(parse.NewTsParser()),
	).RegisterPostProcessor().Render()

	// f, _ := os.Open(".scalars.txt")
	// pgs.Init(
	// 	pgs.ProtocInput(f),
	// ).RegisterModule(
	// 	gopl.GoGen(),
	// 	ts.TsGen(parse.NewTsParser()),
	// ).RegisterPostProcessor(
	// 	pgsgo.GoFmt(),
	// ).Render()
}
