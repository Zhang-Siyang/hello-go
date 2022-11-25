package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"hello/define"
	"hello/pkg/machine"
)

type CheckController struct{}

func (CheckController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"t": time.Now().Format(time.RFC1123Z)})
	return
}

func (CheckController) Envs(c *gin.Context) {
	c.JSON(http.StatusOK, os.Environ())
}

func (CheckController) MachineInfo(c *gin.Context) {
	hostname, _ := os.Hostname()
	ip := machine.GetIP(c)

	c.Header(define.RespHeaderIP, ip)
	c.JSON(http.StatusOK, gin.H{
		define.RespHeaderHostname: hostname,
		define.RespHeaderIP:       ip,
	})
}
