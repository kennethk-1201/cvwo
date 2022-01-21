package helper

import (
	"log"
	"database/sql"
	"os"
)

func CheckErr(err error) {
	if err != nil {
        log.Fatalf("Error occured: %q", err)
	}
}

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	CheckErr(err)

    return db
}