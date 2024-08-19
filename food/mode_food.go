package food

type Food struct {
	Value float64 `json:"value" sql:"value"`
	Time  uint64  `json:"time" sql:"time_taken"`
}

func (Food) TableName() string {
	return "food"
}

type Dura []Food

func (a Dura) Len() int           { return len(a) }
func (a Dura) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Dura) Less(i, j int) bool { return a[i].Time < a[j].Time }
