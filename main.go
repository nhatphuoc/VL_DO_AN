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
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

	var broker = "localhost"
    var port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID("go_mqtt_client")
    opts.SetUsername("emqx")
    opts.SetPassword("public")
    opts.SetDefaultPublishHandler(mqttServer.MessagePubHandler)
    opts.OnConnect = mqttServer.ConnectHandler
    opts.OnConnectionLost = mqttServer.ConnectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	mqttServer.Sub(client)

    //client.Disconnect(250)
    
    var err error
	if database.DB,err = database.CreateDB(); err != nil {
		fmt.Println(err.Error())
	}

	a := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:3000"}
	a.Use(cors.New(config))
	r := a.Group("/api")
	r1 := r.Group("/schedule")
	{
		r1.POST("", schedule.CreateSchedule(database.DB))
        r1.PATCH(":id", schedule.UpdateSchedule(database.DB))
		r1.GET("", schedule.ListSchedule(database.DB))
		r1.DELETE(":id", schedule.DeleteSchedule(database.DB))
	}
    r2 := r.Group("/environment")
	{
		r2.POST("/create", environment.CreateEnvironment(database.DB))
		r2.POST("/list", environment.ListEnvironment(database.DB))
	}
    r3 := r.Group("/gallery")
	{
		r3.POST("/create", gallery.CreateGallery(database.DB))
		r3.POST("/list", gallery.ListGallery(database.DB))
	}
	r4 := r.Group("/homeData")
	{
		r4.GET("/", home.GetHomeData(database.DB))
	}
	r5 := r.Group("/video")
	{
		r5.POST("/list", video.ListVideo(database.DB))
	}
	r6 := r.Group("/log")
	{
		r6.POST("/list", log.ListLog(database.DB))
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
				"url": "fke_url",
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




	a.Run(":3000")

}
