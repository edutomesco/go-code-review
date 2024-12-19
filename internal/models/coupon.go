package models

import (
	"errors"
	appError "github.com/edutomesco/coupons/internal/models/errors"
)

var ErrInvalidMinBasketValue = appError.ErrInvalidBodyJSON(errors.New("min basket value should be greater or equal than discount"))

type Coupon struct {
	id             string
	discount       int
	code           string
	minBasketValue int
}

func NewCoupon(id string, discount int, code string, minBasketValue int) (Coupon, error) {
	if minBasketValue < discount {
		return Coupon{}, ErrInvalidMinBasketValue
	}

	return Coupon{
		id:             id,
		discount:       discount,
		code:           code,
		minBasketValue: minBasketValue,
	}, nil
}

func (c Coupon) ID() string {
	return c.id
}

func (c Coupon) Discount() int {
	return c.discount
}

func (c Coupon) Code() string {
	return c.code
}

func (c Coupon) MinBasketValue() int {
	return c.minBasketValue
}
