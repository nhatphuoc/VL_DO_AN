package mqttServer

import (
	"encoding/json"
	"fmt"
	"go-module/database"
	"go-module/environment"
	"go-module/gallery"
	"go-module/log"
	"go-module/schedule"
	"go-module/timeeat"
	"go-module/video"
	"go-module/water"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Sub(client mqtt.Client) {
	topics := []string{"sensor_state", "add_image", "add_video", "request_feed_time", "dev_info", "log", "time_eat", "water_added"}
	for _, topic := range topics {
		token := client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic: %s", topic)
	}
}

type Sensor_State struct {
	Temperature float64 `json:"temp" `
	Humidity    float64 `json:"humid" `
	Food        float64 `json:"food" `
	Water       float64 `json:"water" `
}

type Image struct {
	Url string `json:"url" sql:"url"`
}

type Dev_Info struct {
	Software string `json:"software"`
	Ip       string `json:"ip"`
	Board    string `json:"board"`
	Wifi     string `json:"wifi"`
}

var DevInfo Dev_Info
var DeviceInfomation *Dev_Info = nil
var HomeData Sensor_State

func Received_Dev_Info(payload []byte) {
	json.Unmarshal(payload, &DevInfo)

	DeviceInfomation = &DevInfo
}

func Reiceve_Sensor_State(payload []byte) {
	err := json.Unmarshal(payload, &HomeData)
	if err != nil {
		fmt.Println("Line 61", err.Error())
	}
	nowsub := time.Now()
	now := nowsub.Unix()
	exec := fmt.Sprintf(`insert into %s (temperature, humidity, food, water, time_taken)
	values (?,?,?,?,?)`, environment.Environment{}.TableName())
	_, err = database.DB.Exec(exec, HomeData.Temperature, HomeData.Humidity, HomeData.Food, HomeData.Water, now)

	if err != nil {
		fmt.Println("Reiceve_Sensor_State Sensor error:", err)
	}
}

func Reiceve_image(payload []byte) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	url := string(payload)

	exec := fmt.Sprintf(`insert into %s (url,time_taken)
	values (?,?)`, gallery.Gallery{}.TableName())
	_, err := database.DB.Exec(exec, url, unixTimestamp)

	if err != nil {
		fmt.Println("Receive_image error:", err)
	}
}

func Reiceve_food(payload []byte) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	var ss timeeat.Timeeat
	json.Unmarshal(payload, &ss)

	exec := fmt.Sprintf(`insert into %s (food,time_taken)
	values (?,?)`, timeeat.Timeeat{}.TableName())
	_, err := database.DB.Exec(exec, ss.Value, unixTimestamp)

	if err != nil {
		fmt.Println("Receive_food error:", err)
	}
}

func Reiceve_water(payload []byte) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	var ss water.Water
	json.Unmarshal(payload, &ss)

	exec := fmt.Sprintf(`insert into %s (water,time_taken)
	values (?,?)`, water.Water{}.TableName())
	_, err := database.DB.Exec(exec, ss.Value, unixTimestamp)

	if err != nil {
		fmt.Println("Reiceve_water error:", err)
	}
}

func Reiceve_video(payload []byte) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	url := string(payload)

	exec := fmt.Sprintf(`insert into %s (url,time_taken)
	values (?,?)`, video.Video{}.TableName())
	_, err := database.DB.Exec(exec, url, unixTimestamp)

	if err != nil {
		fmt.Println("Reiceve_video error:", err)
	}
}

func Reiceve_log(payload []byte) {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	url := string(payload)

	exec := fmt.Sprintf(`insert into %s (url,time_taken)
	values (?,?)`, log.Log{}.TableName())
	_, err := database.DB.Exec(exec, url, unixTimestamp)

	if err != nil {
		fmt.Println("Reiceve_log error:", err)
	}
}

func Write_feed_time(client mqtt.Client) {
	query := fmt.Sprintf(`SELECT id, feed_value, feed_time, feed_duration, url, isOn FROM %s WHERE isOn = 1`, schedule.Schedule{}.TableName())
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Write_feed_time error:", err)
	}
	var sche schedule.Schedule
	listSche := make([]schedule.Schedule, 0)
	for rows.Next() {
		err = rows.Scan(&sche.ID, &sche.Value, &sche.Time, &sche.Feed_Duration, &sche.Url, &sche.IsOn)
		if err != nil {
			fmt.Println("Write_feed_time error:", err)
		}
		if err != nil {
			fmt.Println("Write_feed_time error:", err)
		}
		listSche = append(listSche, sche)
	}
	var payload string
	for _, schedu := range listSche {
		tmp := strings.Split(schedu.Time, ":")
		hours, err := strconv.ParseInt(tmp[0], 10, 8)
		if err != nil {
			hours = 25
		}
		minutes, err := strconv.ParseInt(tmp[1], 10, 8)
		if err != nil {
			minutes = 61
		}
		payload += fmt.Sprintf("{%d %d fake_url %d %d}", hours, minutes, schedu.Value, sche.Feed_Duration)
	}
	client.Publish("write_feed_time", 2, false, payload)

}

func Write_Feed_Now(client mqtt.Client) {
	client.Publish("feed_now", 2, false, "1")
}

func Write_Restart(client mqtt.Client) {
	client.Publish("restart", 2, false, "1")
}

func Write_Callfunc(client mqtt.Client) {
	client.Publish("call", 2, false, "fake_url")
}

func Write_DevInfo(client mqtt.Client) {
	token := client.Publish("get_dev_info", 2, false, "1")
	token.WaitTimeout(time.Second)
}
