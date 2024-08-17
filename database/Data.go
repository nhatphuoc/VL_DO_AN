package database

import "database/sql"

var DB *sql.DB

func CreateDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "gotest:gotest@tcp(db:3306)/db")
	if err != nil {
		return nil, err
	}
	// db, err := sql.Open("mysql", "nhatphuoc:123456789@tcp(localhost:3306)/")
	// if err != nil {
	//     return nil,err
	// }
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS DB;")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("USE DB;")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Schedule(
        id INT PRIMARY KEY AUTO_INCREMENT,
        feed_value INT NOT NULL,
        feed_time TIME NOT NULL,
        isOn BOOL
    );`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS Environment(
		temperature int not null,
		humidity int not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS Log(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS Gallery(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}


	_, err = db.Exec(`create table IF NOT EXISTS Video(
		url longtext not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS Food(
		food int not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table IF NOT EXISTS Water(
		water int not null,
		time_taken int UNSIGNED not null
	);`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
