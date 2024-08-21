package timeeat

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

func ListTimeeat(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Timeeat{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var fd Timeeat
		var listFD []Timeeat
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

		sub := make([]Timeeat, 0)
		for _, Food1 := range listFD {
			timeStamp := time.Unix(int64(Food1.Time), 0)
			startOfDay := time.Date(timeStamp.Year(), timeStamp.Month(), timeStamp.Day(), 0, 0, 0, 0, timeStamp.Location())
			startOfDayUnix := startOfDay.Unix()
			Food1.Time = uint64(startOfDayUnix)
			sub = append(sub, Food1)
		}

		groupedData := make(map[uint64][]Timeeat)
		for _, d := range sub {
			groupedData[d.Time] = append(groupedData[d.Time], d)
		}

		re := make([]Timeeat, 0)
		for timestamp, food2 := range groupedData {
			var sumValue float64
			for _, food3 := range food2 {
				sumValue += float64(food3.Value)
			}
			avgValue := sumValue / float64(len(food2))

			re = append(re, Timeeat{
				Value: avgValue,
				Time:  timestamp,
			})
		}
		sort.Sort(Dura(re))

		c.JSON(http.StatusOK, gin.H{
			"feedList": re,
		})
	}
}
