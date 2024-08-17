package food

type Food struct {
	Value int    `json:"value" sql:"value"`
	Time  uint64 `json:"time" sql:"time_taken"`
}

func (Food) TableName() string {
	return "food"
}
