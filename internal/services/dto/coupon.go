package dto

type CreateCouponRequest struct {
	Discount       int
	Code           string
	MinBasketValue int
}

type ApplyCouponRequest struct {
	Code   string
	Basket BasketRequest
}

type BasketRequest struct {
	Value int
}
