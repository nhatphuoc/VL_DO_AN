package mqttServer

import (
	"encoding/json"
	"fmt"
	"go-module/database"
	"go-module/food"
	"go-module/gallery"
	"go-module/log"
	"go-module/schedule"
	"go-module/video"
	"go-module/water"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Sub(client mqtt.Client){
	topics := []string{"device_state","add_image","add_video", "request_feed_time", "dev_info", "log","time_eat","water_added"}
	for _, topic := range topics {
		token := client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic: %s", topic)
	}
}

type Sensor_State struct {
	Temperature int `json:"temp" `
	Humidity    int `json:"humid" `
	Food  int    `json:"food" `
	Water int    `json:"water" `
}

type Image struct {
	Url  string `json:"url" sql:"url"`
}

type Dev_Info struct {
	Software string `json:"software"`
	Ip string `json:"ip"`
	Board string `json:"board"`
	Wifi string `json:"wifi"`
}

var DeviceInfomation *Dev_Info = nil
var HomeData Sensor_State

func Received_Dev_Info( payload []byte) {
	var ss Dev_Info
	json.Unmarshal(payload, &ss)

	DeviceInfomation = &ss
}

func Reiceve_Sensor_State(payload []byte) {
	json.Unmarshal(payload, &HomeData)
}

func Reiceve_image(payload []byte ) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	url := string(payload)

	exec := fmt.Sprintf(`insert into %s (url,time_taken)
	value (?,?)`, gallery.Gallery{}.TableName())
	_,err := database.DB.Exec(exec, url, unixTimestamp)

	if err != nil {
		fmt.Println("Receive_image error:",err)
	}
}

func Reiceve_food(payload []byte ) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	var ss food.Food
	json.Unmarshal(payload, &ss)

	exec := fmt.Sprintf(`insert into %s (value,time_taken)
	value (?,?,?)`,food.Food{}.TableName())
	_,err := database.DB.Exec(exec, ss.Value, unixTimestamp)

	if err != nil {
		fmt.Println("Receive_food error:",err)
	}
}

func Reiceve_water(payload []byte ) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	var ss water.Water
	json.Unmarshal(payload, &ss)

	exec := fmt.Sprintf(`insert into %s (value,time_taken)
	value (?,?,?)`,water.Water{}.TableName())
	_,err := database.DB.Exec(exec, ss.Value, unixTimestamp)

	if err != nil {
		fmt.Println("Reiceve_water error:",err)
	}
}

func Reiceve_video(payload []byte )  {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	url := string(payload)

	exec := fmt.Sprintf(`insert into %s (url,time_taken)
	value (?,?)`, video.Video{}.TableName())
	_,err := database.DB.Exec(exec, url, unixTimestamp)

	if err != nil {
		fmt.Println("Reiceve_video error:",err)
	}
}

func Reiceve_log(payload []byte )  {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	url := string(payload)

	exec := fmt.Sprintf(`insert into %s (url,time_taken)
	value (?,?)`, log.Log{}.TableName())
	_,err := database.DB.Exec(exec, url, unixTimestamp)

	if err != nil {
		fmt.Println("Reiceve_log error:",err)
	}
}

func Write_feed_time(client mqtt.Client)  {
	query := fmt.Sprintf(`SELECT * from db.%s`, schedule.Schedule{}.TableName())
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Write_feed_time error:",err)
	}
	var sche schedule.Schedule
	var listSche []schedule.Schedule
	for rows.Next() {
		var t string
		err = rows.Scan(&sche.ID, &sche.Value, &t, &sche.IsOn)
		a := strings.Split(t,":")
		if err != nil {
			fmt.Println("Write_feed_time error:",err)
		}
		sche.Time, err = time.ParseDuration(a[0]+"h"+a[1]+"m"+a[2]+"s")
		if err != nil {
			fmt.Println("Write_feed_time error:",err)
		}
		listSche = append(listSche, sche)
	}
	var payload string
	for _, schedu := range listSche {
		v := schedu.IsOn
		hours := schedu.Time / time.Hour
		minutes := (schedu.Time % time.Hour) / time.Minute
		if bool(v) {
			payload += fmt.Sprintf("{%d %d fake_url %d %d}",hours, minutes, schedu.Value, sche.Feed_Duration)
		}
	}
	token := client.Publish("write_feed_time", 0, false, payload)
    token.Wait()

}

func Write_Feed_Now(client mqtt.Client)  {
	token := client.Publish("feed_now", 0, false, 1)
    token.Wait()
}

func Write_Restart(client mqtt.Client)  {
	token := client.Publish("restart", 0, false, nil)
    token.Wait()
}

func Write_Callfunc(client mqtt.Client)  {
	token := client.Publish("call", 0, false, "fake_url")
    token.Wait()
}

