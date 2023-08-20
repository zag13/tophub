package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/api/controller"
	"github.com/zag13/tophub/server/dal/query"
)

func NewFeedRouter(q *query.Query, group *gin.RouterGroup) {
	c := controller.FeedController{
		Q: q,
	}
	group.GET("/data", c.Feed)
}
