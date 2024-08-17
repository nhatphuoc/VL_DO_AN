package water

type Water struct {
	Value int    `json:"value" sql:"value"`
	Time  uint64 `json:"time" sql:"time_taken"`
}

func (Water) TableName() string {
	return "water"
}
