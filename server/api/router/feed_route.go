package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/api/controller"
	"github.com/zag13/tophub/server/dal"
)

func NewFeedRoute(q *dal.Query, group *gin.RouterGroup) {
	c := controller.FeedController{
		TopsModel: dal.NewTopsModel(q),
	}
	group.POST("/feed", c.Feed)
}
