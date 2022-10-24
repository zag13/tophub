package data

import "time"

// Data 应该和task共同依赖于一个数据服务
type Data struct {
	Id         int64
	Tab        string
	Position   uint8
	Title      string
	Url        string
	Extra      string
	SpiderTime time.Time `gorm:"<-:false"`
}

func (Data) TableName() string {
	return "data"
}
