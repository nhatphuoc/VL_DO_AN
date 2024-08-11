package food_drink

import (
	"database/sql"
	"fmt"
	"go-module/common"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateFoodDrink(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context){
		var fd Food_Drink
		if err := c.ShouldBind(&fd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		exec := fmt.Sprintf(`insert into %s (food,drink,time_taken)
		value (?,?,?)`, Food_Drink{}.TableName())
		_,err := db.Exec(exec, fd.Food, fd.Drink, fd.Time)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,true)
	}
}

func ListFoodDrink(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reDate common.Star_End_Day

		if err := c.ShouldBind(&reDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}


		query := fmt.Sprintf(`SELECT * from %s  where time_taken between %d and %d`, Food_Drink{}.TableName(), reDate.StartDate, reDate.EndDate)
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var fd Food_Drink
		var listFD []Food_Drink
		for rows.Next() {
			err = rows.Scan(&fd.Food, &fd.Drink, &fd.Time)

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