package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zag13/tophub/server/dal"
	"github.com/zag13/tophub/server/domain"
)

type FeedController struct {
	NewsModel dal.NewsModel
}

func (fc *FeedController) Feed(c *gin.Context) {
	var req domain.FeedRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Error(domain.ErrorCodeInvalidParameter, err.Error()))
		return
	}

	vals, err := fc.NewsModel.FindLatest(c, map[string]any{"Site": req.Site})
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Error(domain.ErrorCodeInternalError, err.Error()))
		return
	}

	var feeds []domain.Feed
	for _, val := range vals {
		feeds = append(feeds, domain.Feed{
			Site:    val.Site,
			Ranking: val.Ranking,
			Title:   val.Title,
			URL:     val.URL,
		})
	}

	domain.ToSuccessResponse(c, domain.FeedData{List: feeds})
}
