package utils

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var lock = &sync.Mutex{}

var CustomLogger *zap.SugaredLogger

func GetLogger(instantiate func(options ...zap.Option) (*zap.Logger, error)) *zap.SugaredLogger {
	if CustomLogger == nil {
		lock.Lock()
		defer lock.Unlock()
		if CustomLogger == nil {
			logger, err := instantiate()
			if err != nil {
				fmt.Println(err)
			}
			CustomLogger = logger.Sugar()
		}
	}

	return CustomLogger
}
