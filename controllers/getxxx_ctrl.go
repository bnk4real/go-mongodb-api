package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *customersController) GetXXX(c *gin.Context) {

	ctx := context.Context(c)

	resp, err := ctrl.srv.GetXXX(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)
}
