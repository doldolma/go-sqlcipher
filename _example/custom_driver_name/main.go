package main

import (
	"database/sql"

	_ "github.com/doldolma/go-sqlcipher/v4"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
