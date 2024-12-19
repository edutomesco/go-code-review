package server

import (
	ginServer "github.com/edutomesco/coupons/cmd/api/server/gin"
	"github.com/edutomesco/coupons/internal/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func New(name string, engine *ginServer.GinEngine, hc *controllers.HealthController, cc *controllers.CouponController) (*ginServer.GinRuntime, error) {
	rt := ginServer.NewGinRuntime(name, engine, 5*time.Second)

	rt.SetHandlers(func(e *gin.Engine) {
		e.GET("/health", hc.Health())

		v1 := e.Group("/v1")

		coupons := v1.Group("/coupons")

		coupons.POST("", cc.CreateCoupon())
		coupons.GET("/codes", cc.GetCouponCodes())
		coupons.POST("/application", cc.ApplyCoupon())
	})

	return rt, nil
}
