package main

import (
	"fmt"
	"os"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	fmt.Printf("Connecting as %s with password %s\n", dbUser, dbPass)
}
