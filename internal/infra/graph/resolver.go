package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"clean-arch/internal/infra/database"
)

type Resolver struct {
	OrderRepo *database.OrderRepository
}
