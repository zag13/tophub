package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/api/controller"
)

func NewPingRoute(group *gin.RouterGroup) {
	c := controller.PingController{}
	group.GET("/ping", c.Ping)
}
