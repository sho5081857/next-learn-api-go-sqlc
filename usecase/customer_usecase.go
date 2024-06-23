package usecase

import (
	"context"
	"next-learn-go-sqlc/db/sqlc"
)

type ICustomerUsecase interface {
	GetAllCustomers() ([]sqlc.GetAllCustomersRow, error)
	GetFilteredCustomers(query string) ([]sqlc.GetFilteredCustomersRow, error)
	GetCustomerCount() (int64, error)
}

type customerUsecase struct {
	cq sqlc.Querier
}

func NewCustomerUsecase(cq sqlc.Querier) ICustomerUsecase {
	return &customerUsecase{cq}
}

func (cu *customerUsecase) GetAllCustomers() ([]sqlc.GetAllCustomersRow, error) {
	customers, err := cu.cq.GetAllCustomers(context.Background())
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cu *customerUsecase) GetFilteredCustomers(query string) ([]sqlc.GetFilteredCustomersRow, error) {
	customers, err := cu.cq.GetFilteredCustomers(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cu *customerUsecase) GetCustomerCount() (int64, error) {
	count, err := cu.cq.GetCustomerCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}
