package video

type Video struct {
	Url  string `json:"url" sql:"url"`
	Time uint64 `json:"time" sql:"time_taken"`
}

func (Video) TableName() string {
	return "video"
}
