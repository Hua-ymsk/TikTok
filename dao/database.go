package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func Connect() (err error) {
	defer Close()
	db, err = sql.Open("mysql", "tiktok_user:tiktok_passwd_2024@tcp(101.33.204.176:3306)/TikTok?charset=utf8")
	if err != nil {
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return
}

func Close() {
	err := db.Close()
	if err != nil {
		return
	}
}
