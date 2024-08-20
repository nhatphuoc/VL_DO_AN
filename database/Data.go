package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func CreateDB() (*sql.DB, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	print("Done connecting to db")
	// db, err := sql.Open("mysql", "nhatphuoc:123456789@tcp(localhost:3306)/")
	// if err != nil {
	//     return nil,err
	// }

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS schedule(
        id INT PRIMARY KEY AUTO_INCREMENT,
        feed_value INT NOT NULL,
        feed_time TIME NOT NULL,
		feed_duration INT NOT NULL,
		url longtext not null,
        isOn BOOL
    );`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS environment(
		temperature int not null,
		humidity int not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS log(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS gallery(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS video(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS food(
		food int not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS water(
		water int not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
