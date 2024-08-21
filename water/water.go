package water

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

func ListWater(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Water{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var fd Water
		listFD := make([]Water, 0)
		for rows.Next() {
			err = rows.Scan(&fd.Value, &fd.Time)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			listFD = append(listFD, fd)
		}
		sub := make([]Water, 0)
		for _, Water1 := range listFD {
			timeStamp := time.Unix(int64(Water1.Time), 0)
			startOfDay := time.Date(timeStamp.Year(), timeStamp.Month(), timeStamp.Day(), 0, 0, 0, 0, timeStamp.Location())
			startOfDayUnix := startOfDay.Unix()
			Water1.Time = uint64(startOfDayUnix)
			sub = append(sub, Water1)
		}

		groupedData := make(map[uint64][]Water)
		for _, d := range sub {
			groupedData[d.Time] = append(groupedData[d.Time], d)
		}

		re := make([]Water, 0)
		for timestamp, water2 := range groupedData {
			var sumValue float64
			for _, water3 := range water2 {
				sumValue += float64(water3.Value)
			}
			avgValue := sumValue / float64(len(water2))

			re = append(re, Water{
				Value: avgValue,
				Time:  timestamp,
			})
		}
		sort.Sort(Dura(re))

		c.JSON(http.StatusOK, gin.H{
			"waterList": re,
		})
	}
}
