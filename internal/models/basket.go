package models

import (
	"errors"
	appError "github.com/edutomesco/coupons/internal/models/errors"
	_ "github.com/gin-gonic/gin"
)

var ErrInsufficientBasketValue = appError.ErrInvalidBodyJSON(errors.New("insufficient basket value to apply discount"))

type Basket struct {
	value                 int
	appliedDiscount       int
	applicationSuccessful bool
}

func NewBasket(value int, coupon Coupon) (Basket, error) {
	if value >= coupon.MinBasketValue() {
		return Basket{
			value:                 value - coupon.Discount(),
			appliedDiscount:       coupon.Discount(),
			applicationSuccessful: true,
		}, nil
	} else {
		return Basket{}, ErrInsufficientBasketValue
	}
}

func (b Basket) Value() int {
	return b.value
}

func (b Basket) AppliedDiscount() int {
	return b.appliedDiscount
}

func (b Basket) ApplicationSuccessful() bool {
	return b.applicationSuccessful
}
