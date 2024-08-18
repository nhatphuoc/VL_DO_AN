package gallery

type Gallery struct {
	Url  string `json:"url" sql:"url"`
	Time uint64 `json:"time" sql:"time_taken"`
}

func (Gallery) TableName() string {
	return "gallery"
}
