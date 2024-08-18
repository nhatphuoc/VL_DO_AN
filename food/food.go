package food

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context){
		var env Food
		if err := c.ShouldBind(&env); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (value,time_taken)
		value (?,?)`, Food{}.TableName())
		_,err := db.Exec(exec, env.Value, env.Time)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,true)
	}
}


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
			"feedList": listFD,
		})
	}
}