package schedule

import (
	"go-module/common"
)

type Schedule struct {
	ID            int              `json:"id" sql:"id"`
	Value         int              `json:"value" sql:"feed_value"`
	Time          string           `json:"time" sql:"feed_time"`
	Feed_Duration int              `json:"feed_duration" sql:"feed_duration"`
	IsOn          common.ModelBool `json:"isOn" sql:"isOn"`
	Url           string           `json:"url" sql:"url"`
}

func (Schedule) TableName() string {
	return "schedule"
}

type ScheduleCreation struct {
	Value         int              `json:"value" sql:"feed_value"`
	Time          string           `json:"time" sql:"feed_time"`
	Feed_Duration int              `json:"feed_duration" sql:"feed_duration"`
	Url           string           `json:"url" sql:"url"`
	IsOn          common.ModelBool `json:"isOn" sql:"isOn"`
}

type Dura []Schedule

func (a Dura) Len() int           { return len(a) }
func (a Dura) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Dura) Less(i, j int) bool { return a[i].Time < a[j].Time }
