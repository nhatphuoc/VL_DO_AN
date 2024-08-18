package log

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListLog(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := fmt.Sprintf(`SELECT * from %s ORDER BY time_taken DESC LIMIT 100`, Log{}.TableName())
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
			listGal = append(listGal, gal.Url)
		}

		c.JSON(http.StatusOK, gin.H{
			"logs": listGal,
		})
	}
}
