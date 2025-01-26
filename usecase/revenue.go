package usecase

import (
	"context"
	"next-learn-go-sqlc/infrastructure/database/sqlc"
)

type RevenueUseCase interface {
	GetAllRevenues() ([]sqlc.Revenue, error)
}

type revenueUseCase struct {
	rq sqlc.Querier
}

func NewRevenueUseCase(rq sqlc.Querier) RevenueUseCase {
	return &revenueUseCase{rq}
}

func (ru *revenueUseCase) GetAllRevenues() ([]sqlc.Revenue, error) {
	revenues, err := ru.rq.GetAllRevenues(context.Background())
	if err != nil {
		return nil, err
	}
	return revenues, nil
}
