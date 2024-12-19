package interfaces

import (
	"github.com/edutomesco/coupons/internal/models"
)

type CouponRepository interface {
	Save(coupon models.Coupon) error

	GetByCode(code string) (models.Coupon, error)
}
