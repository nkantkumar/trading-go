package main

import (
	"database/sql"
	"log"
	"net/http"
	"trading-go/api"
	"trading-go/internal/db"
	"trading-go/internal/order"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://user:password@localhost/tradingdb?sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	repo := db.NewPostgresRepository(database)
	orderService := order.NewService(repo)
	handler := api.NewHandler(orderService)

	http.HandleFunc("/orders", handler.PlaceOrder)
	http.HandleFunc("/orders/list", handler.ListOrders)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
