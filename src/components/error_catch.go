package components

import (
	"go.uber.org/zap"
)

func CatchError() {

	if err := recover(); err != nil {

		logger.EngineLogger.Error("服务错误信息", zap.Any("errorDetail", err))
	}

}

func CatchPanic() {

	if err := recover(); err != nil {

		logger.EngineLogger.Panic("服务严重异常信息", zap.Any("panicDetail", err))
	}

}
