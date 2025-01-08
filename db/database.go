package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func OpenDatabase() error {

	var err error

	DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("connected to database")
	return nil
}

func CloseDatabase() {
	if DB != nil {
		_ = DB.Close() //veritabanını kapatırken olabilecek hataları _ : ignore
		log.Println("database connection closed")
	}
}
