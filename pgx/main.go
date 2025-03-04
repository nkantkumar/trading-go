package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=nishi password=nishi dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to open connection: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	// Example Query
	var now string
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Printf("Current time in DB: %s\n", now)
}
