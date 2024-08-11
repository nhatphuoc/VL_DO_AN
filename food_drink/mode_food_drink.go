package food_drink

type Food_Drink struct {
	Food  int    `json:"food" sql:"food"`
	Drink int    `json:"drink" sql:"drink"`
	Time  uint64 `json:"time" sql:"time_taken"`
}

func (Food_Drink) TableName() string {
	return "food_drink"
}
