package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"hello/lib/zaps"
)

type TraceIDKeyType string

const TraceIDContextKey = TraceIDKeyType("")

func SetLogger(c *gin.Context) {
	traceID := uuid.NewString()
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), zaps.ContentKey, zaps.Logger(c.Request.Context()).With(zap.String("trace-id", traceID))))
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), TraceIDContextKey, traceID))
}

func GetTraceID(c context.Context) string {
	if s, ok := c.Value(TraceIDContextKey).(string); ok {
		return s
	}
	if ginContext, ok := c.(*gin.Context); ok {
		if s, ok := ginContext.Request.Context().Value(TraceIDContextKey).(string); ok {
			return s
		}
	}
	return "(nil)"
}
