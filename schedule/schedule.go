package schedule

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var sche ScheduleCreation
		if err := c.ShouldBind(&sche); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (feed_value,feed_time,isOn)
		value (?,?,?)`, Schedule{}.tableName())
		_,err := db.Exec(exec, sche.Value, sche.Time, sche.IsOn)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,true)
	}
}

func UpdateSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))

		if err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		return
		}

		var sche ScheduleCreation
		if err := c.ShouldBind(&sche); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`update %s 
		set feed_value=%d,feed_time=%d,isOn=%t
		where id = %d`, Schedule{}.tableName(), sche.Value, sche.Time, sche.IsOn, id)
		_,err = db.Exec(exec)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,true)
	}
}

func DeleteSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		return
		}

		exec := fmt.Sprintf(`delete from %s where id = %d`, Schedule{}.tableName(), id)
		_,err = db.Exec(exec)
		if err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, true)
		return 
	}
}

func ListSchedule(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := fmt.Sprintf(`SELECT * from db.%s`, Schedule{}.tableName())
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var sche Schedule
		var listSche []Schedule
		for rows.Next() {
			var t string
			err = rows.Scan(&sche.ID, &sche.Value, &t, &sche.IsOn)
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
					"error3": err.Error(),
				})
				return 
			}
			listSche = append(listSche, sche)
		}
		
		
		c.JSON(http.StatusOK,gin.H{
			"feedingSchedule": listSche,
		})
	}
}