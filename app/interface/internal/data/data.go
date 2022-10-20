package data

import "time"

// Data 应该和task共同依赖于一个数据服务
type Data struct {
	Id          int64
	Tab         string
	Host        string
	Position    uint8
	Title       string
	Url         string
	Description string
	Image       string
	Extra       string
	SpiderTime  time.Time
}

func (Data) TableName() string {
	return "data"
}
