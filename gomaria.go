package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/alunos")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	now := time.Now()

	fmt.Println("PT BR:", now.Format("02/01/2006"))

}
