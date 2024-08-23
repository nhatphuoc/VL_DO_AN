package video

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVideo(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var gal Video
		if err := c.ShouldBind(&gal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (url,time_taken)
		values (?,?)`, Video{}.TableName())
		_, err := db.Exec(exec, gal.Url, gal.Time)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, true)
	}
}

func ListVideo(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := fmt.Sprintf(`SELECT * from %s`, Video{}.TableName())
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var gal Video
		listGal := make([]Video, 0)
		for rows.Next() {
			err = rows.Scan(&gal.Url, &gal.Time)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			listGal = append(listGal, gal)
		}

		c.JSON(http.StatusOK, gin.H{
			"videos": listGal,
		})
	}
}
