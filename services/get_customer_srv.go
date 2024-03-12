package services

import (
	"context"
)

type GetCustomerResponse struct {
	Data []CustomerData `json:"data"`
}

type CustomerData struct {
	CustomerId   string `json:"customer_id" validate:"required"`
	CustomerName string `json:"customer_name" validate:"required"`
	CustomerType string `json:"customer_type" validate:"required"`
	Remarks      string `json:"Remarks"`
}

func (srv *servicex) GetCustomers(ctx context.Context) (*GetCustomerResponse, error) {

	// query
	getCustomers, err := srv.srv.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}

	var customersData []CustomerData
	if len(getCustomers) > 0 {
		for _, data := range getCustomers {
			customersData = append(customersData, CustomerData{
				CustomerId:   data.CustomerId,
				CustomerName: data.CustomerName,
				CustomerType: data.CustomerType,
				Remarks:      data.Remarks,
			})
		}
	}

	// create a response
	resp := &GetCustomerResponse{
		Data: customersData,
	}

	return resp, nil
}
