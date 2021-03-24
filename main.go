package main

import (
	"os"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	// data, _ := ioutil.ReadAll(os.Stdin)
	// os.WriteFile("./input-kitchent.txt", data, os.ModePerm)

	data, _ := os.Open("./input-kitchent.txt")
	pgs.Init(
		pgs.ProtocInput(data),
	).RegisterModule(
		ASTPrinter(),
		// JSONify(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
