package food

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListFood(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}

		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Food{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var fd Food
		var listFD []Food
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
		
		c.JSON(http.StatusOK,gin.H{
			"foodDrinkHistory": listFD,
		})
	}
}