package main

import (
	"os"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	f, _ := os.Open("./input.txt")
	pgs.Init(
		pgs.ProtocInput(f),
	).RegisterModule(
		ASTPrinter(),
		JSONify(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
