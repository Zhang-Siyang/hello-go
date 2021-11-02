package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemporaryController struct{}

func (TemporaryController) OK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
