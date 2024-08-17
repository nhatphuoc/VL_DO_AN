package log

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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