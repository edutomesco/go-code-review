package controllers

import (
	"github.com/edutomesco/coupons/internal/controllers/datatransfers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, datatransfers.HealthResponse{Success: true})
	}
}
