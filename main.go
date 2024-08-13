package main

import (
	"database/sql"
	"fmt"
	"go-module/environment"
	"go-module/food_drink"
	"go-module/gallery"
	"go-module/home-data"
	"go-module/schedule"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



func createDB() (*sql.DB,error) {
    // db, err := sql.Open("mysql", "gotest:gotest@tcp(db:3306)/db")
	// if err != nil {
	// 	return nil, err
	// }
	db, err := sql.Open("mysql", "nhatphuoc:123456789@tcp(localhost:3306)/")
    if err != nil {
        return nil,err
    }
    _, err = db.Exec("CREATE DATABASE IF NOT EXISTS DB;")
    if err != nil {
        return nil,err
    }

    _, err = db.Exec("USE DB;")
    if err != nil {
        return nil,err
    }

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS Schedule(
        id INT PRIMARY KEY AUTO_INCREMENT,
        feed_value INT NOT NULL,
        feed_time TIME NOT NULL,
        isOn BOOL
    );`)
    if (err != nil) {
        return nil,err
    }

	_, err = db.Exec(`create table IF NOT EXISTS Environment(
		temperature int not null,
		humidity int not null,
		time_taken int UNSIGNED not null
	);`)
    if (err != nil) {
        return nil,err
    }

	_, err = db.Exec(`create table IF NOT EXISTS Gallery(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
    if (err != nil) {
        return nil,err
    }

	_, err = db.Exec(`create table IF NOT EXISTS Food_Drink(
		food int not null,
		drink int not null,
		time_taken int UNSIGNED not null
	);`)
    if (err != nil) {
        return nil,err
    }

    return db,nil
}

func main() {
    var db *sql.DB
    var err error
	if db,err = createDB(); err != nil {
		fmt.Println(err.Error())
	}

	a := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:3000"}
	a.Use(cors.New(config))
	r := a.Group("/api")
	r1 := r.Group("/schedule")
	{
		r1.POST("", schedule.CreateSchedule(db))
        r1.PATCH(":id", schedule.UpdateSchedule(db))
		r1.GET("", schedule.ListSchedule(db))
		r1.DELETE(":id", schedule.DeleteSchedule(db))
	}
    r2 := r.Group("/environment")
	{
		r2.POST("/create", environment.CreateEnvironment(db))
		r2.POST("/list", environment.ListEnvironment(db))
	}
    r3 := r.Group("/gallery")
	{
		r3.POST("/create", gallery.CreateGallery(db))
		r3.POST("/list", gallery.ListGallery(db))
	}
	r4 := r.Group("/foodDrink")
	{
		r4.POST("/create", food_drink.CreateFoodDrink(db))
		r4.POST("/list", food_drink.ListFoodDrink(db))
	}
	r5 := r.Group("/homeData")
	{
		r5.GET("/", home.GetHomeData(db))
	}
	a.Run(":3000")

}
