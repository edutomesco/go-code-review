package interfaces

import (
	"github.com/edutomesco/coupons/internal/models"
	"github.com/edutomesco/coupons/internal/services/dto"
)

type CouponService interface {
	CreateCoupon(req dto.CreateCouponRequest) error
	ApplyCoupon(req dto.ApplyCouponRequest) (models.Basket, error)

	GetCoupons(codes []string) ([]models.Coupon, error)
}
