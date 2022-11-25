package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"hello/define"
	"hello/exec/api/controller"
	"hello/exec/api/middleware"
	"hello/pkg/gin/log"
	"hello/pkg/utils"
	"hello/pkg/zaps"
)

func setRoute(root *gin.Engine) {
	gCheck := root.Group("/check")
	{
		mainCtl := controller.CheckController{}
		gCheck.GET("/machine/info", mainCtl.MachineInfo)
		gCheck.GET("/ping", mainCtl.Ping)
		gCheck.GET("/envs", middleware.UserAuth, mainCtl.Envs)
	}
	tmp := root.Group("/tmp")
	{
		tmpCtl := controller.TemporaryController{}
		tmp.GET("/auth-test", middleware.UserAuth, tmpCtl.OK)
		tmp.GET("/longer/:a/:b", tmpCtl.GetLonger)
	}
}

func getEngine() *gin.Engine {
	r := gin.New()
	r.Use(middleware.SetLogger)
	r.Use(log.LoggerWithConfig(gin.LoggerConfig{
		Formatter: log.LogFormatter,
		Output:    os.Stdout,
		SkipPaths: nil,
	}))
	r.Use(gin.Recovery()) // fallback
	r.Use(log.RecoveryWithZaps)
	r.Use(middleware.HeaderMachine)
	return r
}

func Init() {
	rand.Seed(time.Now().UnixNano())
	time.Local = time.UTC
	_ = os.Setenv("TZ", "UTC")
	middleware.EnableUUIDRandPool()

	zaps.RegisterContextUnboxer(func(ctx context.Context) context.Context {
		if ginContext, ok := ctx.(*gin.Context); ok {
			return ginContext.Request.Context()
		}
		return nil
	})
}

func main() {
	Init()
	logger := zaps.GlobalLogger()
	defer utils.IgnoreErr(logger.Sync)

	go func() {
		logger.Sugar().Info(http.ListenAndServe("localhost:6060", nil))
	}()

	logger.Info(fmt.Sprintf("start, version %s, build time: %d(%v)",
		define.BinaryVersion, define.BinaryBuildTime, time.Unix(define.BinaryBuildTime, 0).Format(time.RFC3339)))

	root := getEngine()

	setRoute(root)

	if err := root.Run(":80"); err != nil {
		logger.Error("fail", zap.Error(err))
	}

}
