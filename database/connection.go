package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	port              = 5432
	user              = "postgres"
	password          = "postgress"
	dbname            = "demo"
	maxOpenConns      = 5
	maxIdleConnection = 15
	connMaxLifeTime   = 1
)

var db *sql.DB

//InitConnection Creates Connection to Database
func InitConnection() error {
	host := "localhost"
	if _, ok := os.LookupEnv("HOST"); ok {
		host = "db"
	}
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConnection)
	db.SetConnMaxLifetime(time.Hour * connMaxLifeTime)
	return nil
}

//GetConnection Returns Connection to Database
func GetConnection() *sql.DB {
	if db == nil {
		InitConnection()
	}
	return db
}
