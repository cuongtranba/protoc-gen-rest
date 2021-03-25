package main

import (
	"os"
	"protoc-gen-rest/plugin"
	"protoc-gen-rest/plugin/ts"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	// pgs.Init(
	// 	pgs.ProtocInput(os.Stdin),
	// ).RegisterModule(
	// 	plugin.ASTPrinter(),
	// 	ts.JSONify(),
	// ).RegisterPostProcessor(
	// 	pgsgo.GoFmt(),
	// ).Render()

	//for debug
	// write data input to file for debug
	// data, _ := ioutil.ReadAll(os.Stdin)
	f, _ := os.Open("./input-kitchent.txt")
	pgs.Init(
		pgs.ProtocInput(f),
	).RegisterModule(
		plugin.ASTPrinter(),
		ts.JSONify(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
