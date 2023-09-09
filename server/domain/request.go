package domain

type Request struct {
	Page     int `json:"page" binding:"gte=1"`
	PageSize int `json:"pageSize" binding:"gte=1,lte=100"`
}
