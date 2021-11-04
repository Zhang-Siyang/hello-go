package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"hello/define"
	"hello/exec/api/controller"
	"hello/exec/api/middleware"
	"hello/pkg/zaps"
)

func setRoute(root *gin.Engine) {
	gCheck := root.Group("/check")
	{
		mainCtl := controller.CheckController{}
		gCheck.GET("/machine/info", mainCtl.MachineInfo)
		gCheck.GET("/ping", mainCtl.Ping)
	}
	tmp := root.Group("/tmp")
	{
		tmpCtl := controller.TemporaryController{}
		tmp.GET("/auth-test", middleware.UserAuth, tmpCtl.OK)
		tmp.GET("/longer/:a/:b", tmpCtl.GetLonger)
	}
}

func getEngine() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.HeaderMachine, middleware.SetLogger)
	return r
}

func main() {
	logger := zaps.GlobalLogger()
	logger.Info(fmt.Sprintf("start, version %s, build time: %d(%v)", define.BinaryVersion, define.BinaryBuildTime, time.Unix(define.BinaryBuildTime, 0).Format(time.RFC3339)))
	root := getEngine()
	gin.DisableConsoleColor()

	setRoute(root)

	if err := root.Run(":80"); err != nil {
		logger.Error("fail", zap.Error(err))
	}

}
