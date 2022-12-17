package config

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=DESKTOP-EN0V4OG;user id=Amirudev;password=amiru")
	if err != nil {
		fmt.Println(err)
	}
	return
}
