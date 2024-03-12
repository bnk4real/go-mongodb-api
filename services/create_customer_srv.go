package services

import (
	"context"
	"go-mongodb-api/repositories"
	"net/http"
)

type CreateCustomerRequest struct {
	CustomerId   string `json:"customer_id" validate:"required"`
	CustomerName string `json:"customer_name" validate:"required"`
	CustomerType string `json:"customer_type" validate:"required"`
	Remarks      string `json:"remarks"`
}

type CreateCustomerResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
}

func (srv *servicex) CreateCustomer(ctx context.Context, req *CreateCustomerRequest) (resp *CreateCustomerResponse, err error) {

	document := repositories.Customers{
		CustomerId:   req.CustomerId,
		CustomerName: req.CustomerName,
		CustomerType: req.CustomerType,
		Remarks:      req.Remarks,
	}

	err = srv.srv.CreateCustomer(ctx, document)
	if err != nil {
		return nil, err
	}

	resp = &CreateCustomerResponse{}
	resp.ResponseStatus.Code = http.StatusOK
	resp.ResponseStatus.Message = "Create customer successfully."

	return resp, nil
}
