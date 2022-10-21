package data

import "time"

type Data struct {
	Id         int64
	Tab        string
	Host       string
	Position   uint8
	Title      string
	Url        string
	Extra      string
	SpiderTime time.Time `gorm:"<-:false"`
}

func (Data) TableName() string {
	return "data"
}
