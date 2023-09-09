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

	vals, err := fc.NewsModel.FindPage(c, map[string]any{}, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Error(domain.ErrorCodeInternalError, err.Error()))
		return
	}

	feeds := []domain.Feed{}
	for i, val := range vals {
		feeds = append(feeds, domain.Feed{
			Site:    val.Site,
			Ranking: val.Ranking,
			Title:   val.Title,
			URL:     val.URL,
		})
		if i >= req.PageSize-1 {
			break
		}
	}

	domain.ToSuccessResponse(c, domain.FeedData{
		List: feeds,
		PageInfo: domain.PageInfo{
			Page:      req.Page,
			PageSize:  req.PageSize,
			TotalPage: domain.GetTotalPage(req.Page, req.PageSize, len(vals)),
		},
	})
}
