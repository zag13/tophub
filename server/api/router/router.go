package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/bootstrap"
)

func Setup(app bootstrap.Application) *gin.Engine {
	r := gin.Default()

	publicRouter := r.Group("")
	NewPingRoute(publicRouter)
	NewFeedRoute(app.Q, publicRouter)

	return r
}
