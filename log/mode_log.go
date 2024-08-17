package log

type Log struct {
	Url  string `json:"url" sql:"url"`
	Time uint64 `json:"time" sql:"time_taken"`
}

func (Log) TableName() string {
	return "log"
}
