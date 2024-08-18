package gallery

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateGallery(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context){
		var gal Gallery
		if err := c.ShouldBind(&gal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (url,time_taken)
		values (?,?)`, Gallery{}.TableName())
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

func ListGallery(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}


		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Gallery{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var gal Gallery
		var listGal []Gallery
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

		c.JSON(http.StatusOK,gin.H{
			"images": listGal,
		})
	}
}