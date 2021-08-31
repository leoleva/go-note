package util

import (
	"database/sql"
	"demoproject/src/config"
	"github.com/go-sql-driver/mysql"
)

func GetDatabase(config config.Database) (*sql.DB, error) {
	cfg := mysql.Config{
		User: config.User,
		Passwd: config.Password,
		Net: config.Net,
		Addr: config.Host,
		DBName: config.Dbname,
		AllowNativePasswords: config.AllowNativePasswords,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return db, err
	}

	err = db.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
}
