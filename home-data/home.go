package home

import (
	"database/sql"
	"fmt"
	"go-module/gallery"
	"go-module/mqttServer"
	"go-module/schedule"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHomeData(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		nowsub := time.Now()
		now := nowsub.Unix()
		startOfDay := time.Date(nowsub.Year(), nowsub.Month(), nowsub.Day(), 0, 0, 0, 0, nowsub.Location())
		
		query := fmt.Sprintf(`SELECT * from db.%s`, schedule.Schedule{}.TableName())
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var sche  schedule.Schedule
		var listSche []schedule.Schedule
		for rows.Next() {
			var t string
			err = rows.Scan(&sche.ID, &sche.Value, &t,&sche.Feed_Duration, &sche.IsOn)
			a := strings.Split(t,":")

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return 
			}
			sche.Time, err = time.ParseDuration(a[0]+"h"+a[1]+"m"+a[2]+"s")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return 
			}
			listSche = append(listSche, sche)
		}

		sort.Sort(schedule.Dura(listSche))
		a:= startOfDay.Add(listSche[len(listSche)-1].Time-24*time.Hour).Unix()   
		fa :=  listSche[len(listSche)-1]
		b:= startOfDay.Add(listSche[len(listSche)-1].Time-24*time.Hour).Unix()  
		fb :=  listSche[len(listSche)-1]

		if now > startOfDay.Add(listSche[len(listSche)-1].Time).Unix() {
			a = startOfDay.Add(listSche[len(listSche)-1].Time).Unix()
			fa =  listSche[len(listSche)-1]
			b = startOfDay.Add(listSche[0].Time+24*time.Hour).Unix()
			fb =  listSche[0]

		} else {
			for _,t := range listSche {
				if now > a && now < b {
					break;
				} else {
					a = b
					b = startOfDay.Add(t.Time).Unix()
					fa = fb
					fb = t
				}
			}
		}

		query = fmt.Sprintf(`SELECT * from %s  ORDER BY time_taken DESC LIMIT 1`, gallery.Gallery{}.TableName())
		row := db.QueryRow(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var gal gallery.Gallery
		err = row.Scan(&gal.Url, &gal.Time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"food": mqttServer.HomeData.Food,
			"water": mqttServer.HomeData.Water,
			"temp": mqttServer.HomeData.Temperature,
			"humid": mqttServer.HomeData.Humidity,
			"nextFeed": gin.H{
				"value": fb.Value,
				"time":  b,
			},
			"prevFeed": gin.H{
				"value": fa.Value,
				"time":  a,
			},
			"lastImg": gal,
		})

	}
}