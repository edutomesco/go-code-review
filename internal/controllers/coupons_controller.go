package controllers

import (
	"errors"
	"github.com/edutomesco/coupons/internal/controllers/datatransfers"
	"github.com/edutomesco/coupons/internal/interfaces"
	appError "github.com/edutomesco/coupons/internal/models/errors"
	"github.com/edutomesco/coupons/internal/services/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CouponController struct {
	couponService interfaces.CouponService
}

func NewCouponController(cs interfaces.CouponService) *CouponController {
	return &CouponController{
		couponService: cs,
	}
}

func (c *CouponController) CreateCoupon() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req := datatransfers.CreateCouponRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.Error(appError.ErrInvalidBodyJSON(err))
			return
		}

		err := c.couponService.CreateCoupon(dto.CreateCouponRequest{
			Discount:       req.Discount,
			Code:           req.Code,
			MinBasketValue: req.MinBasketValue,
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, datatransfers.CreateCouponResponse{Success: true})
	}
}

func (c *CouponController) GetCouponCodes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		codes, ok := ctx.GetQueryArray("code")
		if !ok {
			_ = ctx.Error(appError.ErrInvalidBodyJSON(errors.New("empty coupon codes")))
			return
		}

		res, err := c.couponService.GetCoupons(codes)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, datatransfers.MapToGetCodesCouponResponse(res))
	}
}

func (c *CouponController) ApplyCoupon() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req := datatransfers.ApplyCouponRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			_ = ctx.Error(appError.ErrInvalidBodyJSON(err))
			return
		}

		basket, err := c.couponService.ApplyCoupon(dto.ApplyCouponRequest{
			Code:   req.Code,
			Basket: dto.BasketRequest{Value: req.Basket.Value},
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, datatransfers.MapToApplyCouponResponse(basket))
	}
}
