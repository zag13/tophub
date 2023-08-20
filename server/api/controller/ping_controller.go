package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct {
}

func (c *PingController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
