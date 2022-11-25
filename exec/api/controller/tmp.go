package controller

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"hello/pkg/zaps"
	"hello/protos"
)

type TemporaryController struct{}

var (
	jackieChanClient = struct {
		sync.Once
		client protos.JackieChanClient
	}{}
)

func (TemporaryController) OK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (t TemporaryController) GetLonger(c *gin.Context) {
	a, b := c.Param("a"), c.Param("b")

	chanClient := t.getJackieChanClient(c)
	ctx, cancel := context.WithTimeout(c, time.Second)
	defer cancel()
	longerReply, err := chanClient.GetLonger(ctx, &protos.LongerRequest{
		M: a,
		N: b,
	})
	if err != nil {
		zaps.Logger(c).Fatal(fmt.Sprintf("could not get longer: %v", err))
	}

	c.JSON(http.StatusOK, gin.H{
		"longer": longerReply.Longer,
	})
}

func (TemporaryController) getJackieChanClient(ctx context.Context) protos.JackieChanClient {
	jackieChanClient.Do(func() {
		// TODO 是不是要加一个 panic
		_ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		conn, _ := grpc.DialContext(_ctx, "localhost:3001", grpc.WithInsecure(), grpc.WithBlock())
		client := protos.NewJackieChanClient(conn)
		jackieChanClient.client = client
	})
	return jackieChanClient.client
}
