package schedule

import (
	"go-module/common"
	"time"
)

type Schedule struct {
	ID    int       `json:"id" sql:"id"`
	Value int       `json:"value" sql:"feed_value"`
	Time  time.Duration       `json:"time" sql:"feed_time"`
	IsOn  common.ModelBool `json:"isOn" sql:"isOn"`
}

func (Schedule) tableName() string {
	return "schedule"
}

type ScheduleCreation struct {
	Value int       `json:"value" sql:"feed_value"`
	Time  time.Duration       `json:"time" sql:"feed_time"`
	IsOn  common.ModelBool `json:"isOn" sql:"isOn"`
}