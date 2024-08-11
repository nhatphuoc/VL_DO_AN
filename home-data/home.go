package home

import (
	"database/sql"
	"fmt"
	"go-module/environment"
	"go-module/food_drink"
	"go-module/gallery"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomeData(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var err error
		query := fmt.Sprintf(`SELECT * from %s  ORDER BY time_taken DESC LIMIT 1`, food_drink.Food_Drink{}.TableName())
		row := db.QueryRow(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var fd food_drink.Food_Drink
		err = row.Scan(&fd.Food, &fd.Drink, &fd.Time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}

		query = fmt.Sprintf(`SELECT * from %s  ORDER BY time_taken DESC LIMIT 1`, environment.Enviroment{}.TableName())
		row = db.QueryRow(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var env environment.Enviroment
		err = row.Scan(&env.Temperature, &env.Humidity, &env.Time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		

		query = fmt.Sprintf(`SELECT * from %s  ORDER BY time_taken DESC LIMIT 1`, gallery.Gallery{}.TableName())
		row = db.QueryRow(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		var gal gallery.Gallery
		err = row.Scan(&gal.Url, &gal.Time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}

		c.JSON(http.StatusOK,gin.H{
			"foodRemaining": fd.Food,        
			"waterRemaining": fd.Drink,      
			"roomTemperature": env.Temperature,     
			"roomHumidity": env.Humidity,            
			"latestImage": gal,
		})
	}
}