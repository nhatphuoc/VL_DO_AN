package timeeat

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"

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

		query := fmt.Sprintf(`SELECT * from %s where time_taken between %d and %d`, Timeeat{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var fd Timeeat
		listFD := make([]Timeeat, 0)
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

		c.JSON(http.StatusOK, gin.H{
			"feedList": listFD,
		})
	}
}
