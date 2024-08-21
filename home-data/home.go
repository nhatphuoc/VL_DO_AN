package home

import (
	"database/sql"
	"fmt"
	"go-module/gallery"
	"go-module/mqttServer"
	"go-module/schedule"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHomeData(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := fmt.Sprintf(`SELECT * from %s WHERE isOn = 1 ORDER BY feed_time DESC`, schedule.Schedule{}.TableName())
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  28,
			})
			return
		}
		var sche schedule.Schedule
		var listSche []schedule.Schedule
		for rows.Next() {
			err = rows.Scan(&sche.ID, &sche.Value, &sche.Time, &sche.Feed_Duration, &sche.Url, &sche.IsOn)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"line":  40,
				})
				return
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"line":  50,
				})
				return
			}
			listSche = append(listSche, sche)
		}

		query = fmt.Sprintf(`SELECT * from %s  ORDER BY time_taken DESC LIMIT 1`, gallery.Gallery{}.TableName())
		row := db.QueryRow(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  62,
			})
			return
		}
		var gal gallery.Gallery
		err = row.Scan(&gal.Url, &gal.Time)
		if err != nil {
			gal.Url = ""
		}

		if len(listSche) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"food":  mqttServer.HomeData.Food,
				"water": mqttServer.HomeData.Water,
				"temp":  mqttServer.HomeData.Temperature,
				"humid": mqttServer.HomeData.Humidity,
				"nextFeed": gin.H{
					"value":         -1,
					"time":          "null",
					"feed_duration": -1,
				},
				"prevFeed": gin.H{
					"value":         -1,
					"time":          "null",
					"feed_duration": -1,
				},
				"lastImg": gal,
			})
		} else {

			now := time.Now()
			now_str := now.Format(time.TimeOnly)[:5]
			previous := listSche[0]
			for _, schedule := range listSche {
				if schedule.Time <= now_str && schedule.Time > previous.Time {
					previous = schedule
				}
			}

			next := listSche[0]

			for _, schedule := range listSche {
				sched_time := ""
				if schedule.Time > now_str {
					sched_time = schedule.Time
				} else {
					sched_time = schedule.Time
					sched_time = "3" + sched_time[1:]
				}
				if sched_time < next.Time {
					next = schedule
				}
			}

			next.Time = next.Time[:5]
			next.Time = next.Time[:2] + next.Time[2:]

			previous.Time = previous.Time[:5]
			previous.Time = previous.Time[:2] + previous.Time[2:]

			c.JSON(http.StatusOK, gin.H{
				"food":  mqttServer.HomeData.Food,
				"water": mqttServer.HomeData.Water,
				"temp":  mqttServer.HomeData.Temperature,
				"humid": mqttServer.HomeData.Humidity,
				"nextFeed": gin.H{
					"value":         next.Value,
					"time":          next.Time[:2] + next.Time[3:5],
					"feed_duration": next.Feed_Duration,
				},
				"prevFeed": gin.H{
					"value":         previous.Value,
					"time":          previous.Time[:2] + previous.Time[3:5],
					"feed_duration": previous.Feed_Duration,
				},
				"lastImg": gal,
			})
		}

	}
}
