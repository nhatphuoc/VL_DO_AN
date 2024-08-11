package environment

type Enviroment struct {
	Temperature int    `json:"temperature" sql:"temperature"`
	Humidity    int    `json:"humidity" sql:"humidity"`
	Time        uint64 `json:"time" sql:"time_taken"`
}

func (Enviroment) TableName() string {
	return "environment"
}
