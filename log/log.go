package log

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateLog(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context){
		var gal Log
		if err := c.ShouldBind(&gal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (url,time_taken)
		value (?,?)`, Log{}.TableName())
		_,err := db.Exec(exec, gal.Url, gal.Time)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,true)
	}
}

func ListLog(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := fmt.Sprintf(`SELECT * from %s`, Log{}.TableName())
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var gal Log
		var listGal []string
		for rows.Next() {
			err = rows.Scan(&gal.Url, &gal.Time)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			timestamp := time.Unix(int64(gal.Time), 0)
			a := fmt.Sprintf("%v", timestamp)
			str := a + ": " + gal.Url
			listGal = append(listGal,str)
		}

		c.JSON(http.StatusOK, gin.H{
			"logs": listGal,
		})
	}
}