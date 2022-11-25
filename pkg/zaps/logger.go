package zaps

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_logger        *zap.Logger
	loggerInitOnce sync.Once
	contextUnboxer []func(context.Context) context.Context
)

type ContentKeyType string

const ContentKey = ContentKeyType("")

func Logger(c context.Context) *zap.Logger {
	if v, ok := c.Value(ContentKey).(*zap.Logger); ok {
		return v
	}
	for _, f := range contextUnboxer {
		if ctx := f(c); ctx != nil {
			if v, ok := ctx.Value(ContentKey).(*zap.Logger); ok {
				return v
			}
		}
	}
	return GlobalLogger()
}

// RegisterContextUnboxer 注册 Context 的拆箱器，比如 1.8 版本之前的 gin.Context 和 底层的 c.Request.Context() 不相通
func RegisterContextUnboxer(f func(context.Context) context.Context) {
	contextUnboxer = append(contextUnboxer, f)
	return
}

func GlobalLogger() *zap.Logger {
	loggerInitOnce.Do(
		func() {
			encoderConfig := zap.NewProductionEncoderConfig()
			encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), os.Stdout, zap.DebugLevel)
			_logger = zap.New(core)
		},
	)
	return _logger
}
