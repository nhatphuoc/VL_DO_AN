package main

import (
	"fmt"
	"go-module/database"
	"go-module/environment"
	"go-module/gallery"
	"go-module/home-data"
	"go-module/log"
	"go-module/mqttServer"
	"go-module/schedule"
	"go-module/video"
	"go-module/water"
	"go-module/food"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
    var err error
	if database.DB,err = database.CreateDB(); err != nil {
		fmt.Println(err.Error())
	}

	var broker = "comqtt"
    var port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID("go_mqtt_client")
    opts.SetDefaultPublishHandler(mqttServer.MessagePubHandler)
    opts.OnConnect = mqttServer.ConnectHandler
    opts.OnConnectionLost = mqttServer.ConnectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	mqttServer.Sub(client)

    //client.Disconnect(250)

	a := gin.Default()
	a.Use(cors.Default())
	a.Use(static.Serve("/", static.LocalFile("./static", true)))

	r := a.Group("/api")

	r1 := r.Group("/schedule")
	{
		r1.POST("/", schedule.CreateSchedule(database.DB))
        r1.PATCH(":id", schedule.UpdateSchedule(database.DB))
		r1.GET("/", schedule.ListSchedule(database.DB))
		r1.DELETE(":id", schedule.DeleteSchedule(database.DB))
	}

    r2 := r.Group("/environment")
	{
		r2.POST("/create", environment.CreateEnvironment(database.DB))
		r2.POST("/list", environment.ListEnvironment(database.DB))
	}

    r3 := r.Group("/image")
	{
		r3.POST("/", gallery.CreateGallery(database.DB))
		r3.GET("/", gallery.ListGallery(database.DB))
	}

	r4 := r.Group("/homeData")
	{
		r4.GET("/", home.GetHomeData(database.DB))
	}

	r5 := r.Group("/video")
	{
		r5.POST("/", video.ListVideo(database.DB))
	}

	r6 := r.Group("/log")
	{
		r6.POST("/", log.ListLog(database.DB))
	}

	r7 := r.Group("/status")
	{
		r7.GET("/",func (c*gin.Context){
			token := client.Publish("get_dev_info", 0, false, nil)
			token.Wait()
			for mqttServer.DeviceInfomation == nil {
			}
			c.JSON(http.StatusOK,mqttServer.DeviceInfomation)
			mqttServer.DeviceInfomation = nil
		})
	}

	r8 := r.Group("/camera")
	{
		r8.GET("/", func (c*gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"url": "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
			})
		})
	}

	r9 := r.Group("/feedNow")
	{
		r9.POST("/", func (c*gin.Context){
			mqttServer.Write_Feed_Now(client)
			c.JSON(http.StatusOK,true)
		})
	}

	r10 := r.Group("/restart")
	{
		r10.POST("/", func (c*gin.Context){
			mqttServer.Write_Restart(client)
			c.JSON(http.StatusOK,true)
		})
	}

	r11 := r.Group("/call")
	{
		r11.GET("/", func (c*gin.Context){
			mqttServer.Write_Callfunc(client)
			c.JSON(http.StatusOK,gin.H{
				"url": "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
			})
		})
	}

	r12 := r.Group("/water")
	{
		r12.POST("/", water.ListWater(database.DB))
	}

	r13 := r.Group("/food")
	{
		r13.POST("/", food.ListFood(database.DB))
	}


	a.Run(":3000")

}
