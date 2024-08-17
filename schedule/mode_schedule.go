package schedule

import (
	"go-module/common"
	"time"
)

type Schedule struct {
	ID    int       `json:"id" sql:"id"`
	Value int       `json:"value" sql:"feed_value"`
	Time  time.Duration       `json:"time" sql:"feed_time"`
	Feed_Duration int `json:"feed_duration" sql:"feed_duration"`
	IsOn  common.ModelBool `json:"isOn" sql:"isOn"`
}

func (Schedule) TableName() string {
	return "schedule"
}

type ScheduleCreation struct {
	Value int       `json:"value" sql:"feed_value"`
	Time  time.Duration       `json:"time" sql:"feed_time"`
	Feed_Duration int `json:"feed_duration" sql:"feed_duration"`
	IsOn  common.ModelBool `json:"isOn" sql:"isOn"`
}

type Dura []Schedule
func (a Dura) Len() int           { return len(a) }
func (a Dura) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Dura) Less(i, j int) bool { return int(a[i].Time) < int(a[j].Time) }