package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var EngineLogger *zap.Logger

var PbLogger *zap.Logger

var GinLogger *zap.Logger

func init() {

	EngineLogger = NewLogger("logs/app_engine.log", zapcore.InfoLevel, 128, 0, 0, true, "EngineApp")

	GinLogger = NewLogger("logs/gin_engine.log", zapcore.InfoLevel, 128, 0, 0, true, "EngineGin")

	PbLogger = NewLogger("logs/pblog_engine.log", zapcore.InfoLevel, 128, 0, 0, true, "EnginePblog")

}
