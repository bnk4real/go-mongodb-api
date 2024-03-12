package controllers

import (
	"context"
	"go-mongodb-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *customersController) CreateCustomer(c *gin.Context) {

	ctx := context.Context(c)
	reqx := &services.CreateCustomerRequest{}

	// Bind the JSON data from the request body to the struct
	if err := c.ShouldBindJSON(&reqx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctrl.srv.CreateCustomer(ctx, reqx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)
}
