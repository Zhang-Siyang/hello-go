package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"hello/define"
	"hello/pkg/zaps"
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

func HeaderMachine(c *gin.Context) {
	c.Next()
	hostname, _ := os.Hostname()
	c.Header(define.RespHeaderHostname, hostname)
}

func UserAuth(c *gin.Context) {
	_, _ = define.ReqHeaderUserID, define.ReqHeaderUserToken
	type UserAuth struct {
		UserID    int64  `header:"X-User-Id" binding:"required"`
		UserToken string `header:"X-User-Token" binding:"required"`
	}
	userAuth := new(UserAuth)
	if err := c.ShouldBindHeader(userAuth); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	} else if userAuth.UserToken != strconv.FormatInt(userAuth.UserID, 10)+"-Token" {
		zaps.Logger(c).Debug(fmt.Sprintf("reject user request, can't pass auth, UserID: %d", userAuth.UserID))
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}
