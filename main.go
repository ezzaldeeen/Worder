package main

import (
	"go.uber.org/zap"
	"worder/pkg/generating"
)

func main() {

	logger, _ := zap.NewDevelopment()

	generator := generating.NewTXTFilesGenerator(logger.Sugar(), "/sdfds", "", 0, 0)
	generator.Generate()
}
