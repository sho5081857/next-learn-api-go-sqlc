package usecase

import (
	"context"
	"next-learn-go-sqlc/db/sqlc"
)

type IRevenueUsecase interface {
	GetAllRevenues() ([]sqlc.Revenue, error)
}

type revenueUsecase struct {
	rq sqlc.Querier
}

func NewRevenueUsecase(rq sqlc.Querier) IRevenueUsecase {
	return &revenueUsecase{rq}
}

func (ru *revenueUsecase) GetAllRevenues() ([]sqlc.Revenue, error) {
	revenues, err := ru.rq.GetAllRevenues(context.Background())
	if err != nil {
		return nil, err
	}
	return revenues, nil
}
