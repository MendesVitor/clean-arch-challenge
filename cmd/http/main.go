package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"clean-arch/internal/infra/database"
	usecase "clean-arch/internal/order"
)

func main() {
	connStr := "postgres://user:password@db:5432/orders_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := database.NewOrderRepository(db)
	service := usecase.NewListOrderUseCase(repo)

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		orders, err := service.ListOrders()
		if err != nil {
			http.Error(w, "Failed to list orders", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	})

	log.Println("HTTP server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
