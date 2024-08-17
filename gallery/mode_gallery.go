package gallery

type Gallery struct {
	Url  string `json:"image_url" sql:"url"`
	Time uint64 `json:"time" sql:"time_taken"`
}

func (Gallery) TableName() string {
	return "gallery"
}
