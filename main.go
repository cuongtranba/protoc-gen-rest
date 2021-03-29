package main

import (
	"os"
	"protoc-gen-rest/plugin/gopl"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	//for debug
	// write data input to file for debug
	// tsParser := parse.NewTsParser()
	// data, _ := ioutil.ReadAll(os.Stdin)
	// ioutil.WriteFile(".scalars.txt", data, 0644)

	// pgs.Init(
	// 	pgs.ProtocInput(os.Stdin),
	// ).RegisterModule(
	// 	gopl.GoGen(),
	// ).RegisterPostProcessor().Render()

	f, _ := os.Open(".scalars.txt")
	pgs.Init(
		pgs.ProtocInput(f),
	).RegisterModule(
		gopl.GoGen(),
		// ts.TsGen(tsParser),
	).RegisterPostProcessor().Render()
}
