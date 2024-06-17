package main

import (
	"clean-arch/internal/infra/database"
	"clean-arch/internal/infra/graph"
	"database/sql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@db/orders_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := database.NewOrderRepository(db)
	resolver := &graph.Resolver{OrderRepo: repo}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("GraphQL server running on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
