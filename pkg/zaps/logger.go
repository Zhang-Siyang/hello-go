package zaps

import (
	"context"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_logger        *zap.Logger
	loggerInitOnce sync.Once
)

type ContentKeyType string

const ContentKey = ContentKeyType("")

func Logger(c context.Context) *zap.Logger {
	if v, ok := c.Value(ContentKey).(*zap.Logger); ok {
		return v
	}
	if ginContext, ok := c.(*gin.Context); ok {
		if v, ok := ginContext.Request.Context().Value(ContentKey).(*zap.Logger); ok {
			return v
		}
	}
	return GlobalLogger()
}

func GlobalLogger() *zap.Logger {
	loggerInitOnce.Do(
		func() {
			_logger = zap.New(zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.AddSync(os.Stdout),
				zapcore.DebugLevel,
			))
		},
	)
	return _logger
}
