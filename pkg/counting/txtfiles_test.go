package counting

import (
	"go.uber.org/zap"
	"testing"
	"worder/utils"
)

func BenchmarkCount(b *testing.B) {
	logger := utils.GetLogger(zap.NewDevelopment)
	// todo: rewrite this
	source := "/Users/ezzaldeen/Hands-on/Golang/worder/data"
	counter := NewWordCounter(logger, source)
	for i := 0; i < b.N; i++ {
		counter.Count()
	}
}
