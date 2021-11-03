package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"hello/controller"
	"hello/define"
	"hello/lib/zaps"
	"hello/middleware"
)

func setRoute(root *gin.Engine) {
	gCheck := root.Group("/check")
	{
		mainCtl := controller.CheckController{}
		gCheck.GET("/machine/info", mainCtl.MachineInfo)
		gCheck.GET("/ping", mainCtl.Ping)
	}
	tmp := root.Group("/tmp", middleware.UserAuth)
	{
		tmpCtl := controller.TemporaryController{}
		tmp.GET("/auth-test", tmpCtl.OK)
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
