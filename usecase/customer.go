package usecase

import (
	"context"
	"next-learn-go-sqlc/infrastructure/database/sqlc"
)

type CustomerUseCase interface {
	GetAllCustomers() ([]sqlc.GetAllCustomersRow, error)
	GetFilteredCustomers(query string) ([]sqlc.GetFilteredCustomersRow, error)
	GetCustomerCount() (int64, error)
}

type customerUseCase struct {
	cq sqlc.Querier
}

func NewCustomerUseCase(cq sqlc.Querier) CustomerUseCase {
	return &customerUseCase{cq}
}

func (cu *customerUseCase) GetAllCustomers() ([]sqlc.GetAllCustomersRow, error) {
	customers, err := cu.cq.GetAllCustomers(context.Background())
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cu *customerUseCase) GetFilteredCustomers(query string) ([]sqlc.GetFilteredCustomersRow, error) {
	customers, err := cu.cq.GetFilteredCustomers(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cu *customerUseCase) GetCustomerCount() (int64, error) {
	count, err := cu.cq.GetCustomerCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}
