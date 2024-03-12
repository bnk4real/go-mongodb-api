package services

import (
	"context"
	"net/http"
	"os"
)

type GetCustomerMock struct {
	Data           CustomerData   `json:"data"`
	ResponseStatus ResponseStatus `json:"responseStatus"`
}

// GetXXX implements Servicex.
func (srv *servicex) GetXXX(context.Context) (resp *GetCustomerMock, err error) {

	customer_id := os.Getenv("customerId")
	customer_name := os.Getenv("customerName")
	customer_type := os.Getenv("customerType")
	customer_remark := os.Getenv("remarks")

	resp = &GetCustomerMock{}
	resp.Data.CustomerId = customer_id
	resp.Data.CustomerName = customer_name
	resp.Data.CustomerType = customer_type
	resp.Data.Remarks = customer_remark
	resp.ResponseStatus.Code = http.StatusOK

	return resp, nil
}
