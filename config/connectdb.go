package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func MYSQL() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Load env failed")
	}
	var dsn = fmt.Sprintf("%v:%v@/%v?parseTime=true", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
