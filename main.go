package main

import "worder/pkg/generator"

func main() {
	gen := generator.TxtFileGenerator{
		Size:        2048,
		Unit:        "MB",
		Count:       100,
		Destination: "data",
	}

	err := gen.Generate()
	if err != nil {
		println(err)
	}
}
