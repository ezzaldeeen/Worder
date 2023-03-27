package generating

import (
	"go.uber.org/zap"
	"testing"
	"worder/utils"
)

func BenchmarkTXTFilesGenerator_Generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger := utils.GetLogger(zap.NewDevelopment)
		// todo: rewrite this
		src := "/Users/ezzaldeen/Hands-on/Golang/worder/resources/sample.txt"
		dest := "/Users/ezzaldeen/Hands-on/Golang/worder/data"
		count := 50
		size := 104857600
		generator := NewTXTFilesGenerator(
			logger, src, dest, count, size)
		for i := 0; i < b.N; i++ {
			generator.Generate()
		}
	}
}
