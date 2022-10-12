package data

type Data struct {
	Id          uint64
	Host        string
	Rank        uint8
	Title       string
	Description string
	Image       string
	Url         string
	Extra       string
	SpiderTime  uint32
}

func (Data) TableName() string {
	return "data"
}
