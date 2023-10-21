package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/api/types"
	"github.com/zag13/tophub/server/dal"
)

type FeedController struct {
	TopsModel dal.TopsModel
}

func (fc *FeedController) Feed(c *gin.Context) {
	var req types.FeedRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: types.ErrorCode_INVALID_PARAMETER, Message: err.Error()})
		return
	}

	vals, err := fc.TopsModel.FindLatest(c, map[string]any{"Site": req.Site})
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: types.ErrorCode_INTERNAL_ERROR, Message: err.Error()})
		return
	}

	var tops []*types.Top
	for _, val := range vals {
		tops = append(tops, &types.Top{
			SpiderTime:  val.SpiderTime.Format(time.DateTime),
			Site:        val.Site,
			Rank:        val.Rank,
			Title:       val.Title,
			Url:         val.URL,
			Description: val.Description,
			Extra:       val.Extra,
		})
	}

	c.JSON(http.StatusOK, types.FeedResponse{Data: &types.FeedData{List: tops}})
}
