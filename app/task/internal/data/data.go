package data

type Data struct {
	Id          uint64
	Host        string
	Position    uint8
	Title       string
	Description string
	Url         string
	Image       string
	Extra       string
	SpiderTime  uint32
}

func (Data) TableName() string {
	return "data"
}
