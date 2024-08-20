package environment

type Environment struct {
	Temperature float64 `json:"temperature" sql:"temperature"`
	Humidity    float64 `json:"humidity" sql:"humidity"`
	Time        uint64  `json:"time" sql:"time_taken"`
}

func (Environment) TableName() string {
	return "environment"
}

type Dura []Environment

func (a Dura) Len() int           { return len(a) }
func (a Dura) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Dura) Less(i, j int) bool { return a[i].Time < a[j].Time }
