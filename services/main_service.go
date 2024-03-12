package services

import (
	"context"
	"go-mongodb-api/repositories"
)

const (
	SuccessCode = "SUCCESS-0000"
	ErrorCode   = "ERROR-0000"
)

type ResponseStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Servicex interface {
	GetXXX(context.Context) (*GetCustomerMock, error)
	GetCustomers(context.Context) (*GetCustomerResponse, error)
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error)
}

type servicex struct {
	srv repositories.CustomersRepositories
}

func NewServicex(
	srv repositories.CustomersRepositories,
) Servicex {
	return &servicex{
		srv,
	}
}
