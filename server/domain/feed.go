package domain

type FeedRequest struct {
	Request
	Site string `json:"site" binding:"required"`
}

type FeedData struct {
	List []Feed `json:"list"`
}

type Feed struct {
	Site    string `json:"site"`
	Ranking int32  `json:"ranking"`
	Title   string `json:"title"`
	URL     string `json:"url"`
}
