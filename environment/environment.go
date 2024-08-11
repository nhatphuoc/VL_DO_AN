package environment

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateEnvironment(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context){
		var env Enviroment
		if err := c.ShouldBind(&env); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (temperature,humidity,time_taken)
		value (?,?,?)`, Enviroment{}.TableName())
		_,err := db.Exec(exec, env.Temperature, env.Humidity, env.Time)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,true)
	}
}

func ListEnvironment(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}


		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Enviroment{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var env Enviroment
		var listEnv []Enviroment
		for rows.Next() {
			err = rows.Scan(&env.Temperature, &env.Humidity, &env.Time)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return 
			}
			listEnv = append(listEnv, env)
		}
		
		c.JSON(http.StatusOK,gin.H{
			"environmentHistory": listEnv,
		})
	}
}