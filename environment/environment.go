package environment

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

// func CreateEnvironment(db *sql.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var env Enviroment
// 		if err := c.ShouldBind(&env); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		exec := fmt.Sprintf(`insert into %s (temperature,humidity,time_taken)
// 		values (?,?,?)`, Enviroment{}.TableName())
// 		_, err := db.Exec(exec, env.Temperature, env.Humidity, env.Time)

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, true)
// 	}
// }

func ListEnvironment(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Environment{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var env Environment
		listEnv := make([]Environment, 0)
		for rows.Next() {
			err = rows.Scan(&env.Temperature, &env.Humidity, &env.Time)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			listEnv = append(listEnv, env)
		}
		sub := make([]Environment, 0)
		for _,Env := range sub {
			timeStamp := time.Unix(int64(Env.Time), 0)
			startOfDay := time.Date(timeStamp.Year(), timeStamp.Month(), timeStamp.Day(), 0, 0, 0, 0, timeStamp.Location())
			startOfDayUnix := startOfDay.Unix()
			Env.Time = uint64(startOfDayUnix)
			sub = append(sub, Env)
		}

		groupedData := make(map[uint64][]Environment)
		for _, d := range sub {
			groupedData[d.Time] = append(groupedData[d.Time], d)
		}

		var re []Environment
		for timestamp, envirs := range groupedData {
			var sumTemp, sumHumidity float64
			for _, envir := range envirs {
				sumTemp += float64(envir.Temperature)
				sumHumidity += float64(envir.Humidity)
			}
			avgTemp := sumTemp / float64(len(envirs))
			avgHumidity := sumHumidity / float64(len(envirs))

			re = append(re, Environment{
				Temperature: avgTemp,
				Humidity:    avgHumidity,
				Time : timestamp,
			})
		}

		sort.Sort(Dura(re))


		c.JSON(http.StatusOK, gin.H{
			"environmentHistory": re,
		})
	}
}
