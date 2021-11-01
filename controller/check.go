package controller

import (
	"hello/define"
	"hello/lib/machine"
	"net/http"
	"os"
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

func (CheckController) MachineInfo(c *gin.Context) {
	zaps.Logger(c).Info("check MachineInfo")
	hostname, _ := os.Hostname()
	ip := machine.GetIP(c)

	c.Header(define.RespHeaderHostname, hostname)
	c.Header(define.RespHeaderIP, ip)
	c.JSON(http.StatusOK, gin.H{
		define.RespHeaderHostname: hostname,
		define.RespHeaderIP:       ip,
	})
}
