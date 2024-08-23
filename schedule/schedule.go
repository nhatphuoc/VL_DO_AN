package schedule

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var sche ScheduleCreation
		if err := c.ShouldBind(&sche); err != nil {
			println("Err 19")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  21,
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (feed_value,feed_time,feed_duration,url,isOn)
		values (?,?,?,?,?)`, Schedule{}.TableName())
		_, err := db.Exec(exec, sche.Value, sche.Time, sche.Feed_Duration, sche.Url, sche.IsOn)

		if err != nil {
			println("Err 32")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  32,
			})
			return
		}
		c.JSON(http.StatusOK, true)
	}
}

func UpdateSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  48,
			})
			return
		}

		var sche ScheduleCreation
		if err := c.ShouldBind(&sche); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  57,
			})
			return
		}

		exec := fmt.Sprintf(`UPDATE %s
		SET feed_value=%d, feed_time='%s', feed_duration=%d, url='%s', isOn=%t
		WHERE id = %d`,
			Schedule{}.TableName(), sche.Value, sche.Time, sche.Feed_Duration, sche.Url, sche.IsOn, id)
		_, err = db.Exec(exec)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  70,
			})
			return
		}
		c.JSON(http.StatusOK, true)
	}
}

func DeleteSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  85,
			})
			return
		}

		exec := fmt.Sprintf(`delete from %s where id = %d`, Schedule{}.TableName(), id)
		_, err = db.Exec(exec)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  95,
			})
			return
		}
		c.JSON(http.StatusOK, true)
	}
}

func ListSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := fmt.Sprintf(`SELECT * from db.%s`, Schedule{}.TableName())
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"line":  111,
			})
			return
		}
		var sche Schedule
		listSche := make([]Schedule, 0)
		for rows.Next() {
			err = rows.Scan(&sche.ID, &sche.Value, &sche.Time, &sche.Feed_Duration, &sche.Url, &sche.IsOn)
			sche.Time = sche.Time[:5]
			sche.Time = sche.Time[:2] + sche.Time[3:]

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"line":  125,
				})
				return
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"line":  133,
				})
				return
			}
			listSche = append(listSche, sche)
		}

		sort.Sort(Dura(listSche))
		c.JSON(http.StatusOK, gin.H{
			"schedule": listSche,
		})
	}
}
