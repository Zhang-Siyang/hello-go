package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"hello/lib/zaps"
)

type CheckController struct{}

func (CheckController) Ping(c *gin.Context) {
	zaps.Logger(c).Info("hello!")
	c.JSON(http.StatusOK, gin.H{"t": time.Now().UnixNano()})
	return
}
