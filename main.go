package main

import (
	"os"
	"protoc-gen-rest/plugin/ts"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	//for debug
	// write data input to file for debug
	// data, _ := ioutil.ReadAll(os.Stdin)
	// ioutil.WriteFile("./input-kitchent.txt", data, 0644)

	// pgs.Init(
	// 	pgs.ProtocInput(os.Stdin),
	// ).RegisterModule(
	// 	plugin.ASTPrinter(),
	// 	ts.JSONify(),
	// ).RegisterPostProcessor(
	// 	pgsgo.GoFmt(),
	// ).Render()

	f, _ := os.Open("./input-kitchent.txt")
	pgs.Init(
		pgs.ProtocInput(f),
	).RegisterModule(
		ts.JSONify(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
