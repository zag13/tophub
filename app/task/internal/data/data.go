package data

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
	SpiderTime  int64
}

func (Data) TableName() string {
	return "data"
}
