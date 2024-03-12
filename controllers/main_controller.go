package controllers

import (
	"go-mongodb-api/services"

	"github.com/gin-gonic/gin"
)

type CustomersController interface {
	GetXXX(c gin.Context) error
	GetCustomers(c gin.Context) error
	CreateCustomer(c gin.Context) error
}

type customersController struct {
	srv services.Servicex
}

func NewCustomerController(
	srv services.Servicex,
) customersController {
	return customersController{
		srv: srv,
	}
}
