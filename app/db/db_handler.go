package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/revel/revel"
)

func Connect() *sql.DB {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	connect_string := revel.Config.StringDefault("db.connect", "root:@tcp(localhost:3306)/db_latihan_pbp")

	db, err := sql.Open(driver, connect_string)
	if err != nil {
		log.Fatalf("Error Connecting to Database: %v", err)
	}

	return db
}
