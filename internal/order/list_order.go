package usecase

import (
	"clean-arch/internal/entity"
	"clean-arch/internal/infra/database"
)

type ListOrderUseCase struct {
	Repo *database.OrderRepository
}

func NewListOrderUseCase(repo *database.OrderRepository) *ListOrderUseCase {
	return &ListOrderUseCase{Repo: repo}
}

func (s *ListOrderUseCase) ListOrders() ([]entity.Order, error) {
	return s.Repo.ListOrders()
}
