package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/dal/query"
)

type FeedController struct {
	Q *query.Query
}

func (c *FeedController) Feed(ctx *gin.Context) {
	// TODO
}
