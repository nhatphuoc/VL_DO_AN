package mqttServer

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    switch msg.Topic() {
        case "device_state":
            Reiceve_Sensor_State(msg.Payload())
        case "add_image":
            Reiceve_image(msg.Payload())
        case "add_video":
            Reiceve_video(msg.Payload())
        case "request_feed_time":
            Write_feed_time(client)
        case "log":
            Reiceve_log(msg.Payload())
        case "time_eat":
            Reiceve_food(msg.Payload())
        case "water_added":
            Reiceve_water(msg.Payload())
        case "dev_info":
            Received_Dev_Info(msg.Payload())
    }
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}



