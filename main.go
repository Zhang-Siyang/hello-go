package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"hello/controller"
	"hello/lib/zaps"
	"hello/middleware"
)

func setRoute(root *gin.Engine) {
	gCheck := root.Group("/check")
	{
		mainCtl := controller.CheckController{}
		gCheck.GET("/ping", mainCtl.Ping)
	}
}

func getEngine() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetLogger)
	return r
}

func main() {
	logger := zaps.GlobalLogger()
	root := getEngine()
	gin.DisableConsoleColor()

	setRoute(root)

	if err := root.Run(":8080"); err != nil {
		logger.Error("fail", zap.Error(err))
	}

}
