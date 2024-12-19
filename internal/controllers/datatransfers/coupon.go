package datatransfers

import "github.com/edutomesco/coupons/internal/models"

type CreateCouponRequest struct {
	Discount       int    `json:"discount" binding:"required"`
	Code           string `json:"code" binding:"required"`
	MinBasketValue int    `json:"min_basket_value" binding:"required"`
}

type ApplyCouponRequest struct {
	Code   string        `json:"code" binding:"required"`
	Basket BasketRequest `json:"basket" binding:"required"`
}

type BasketRequest struct {
	Value int `json:"value" binding:"required"`
}

type CreateCouponResponse struct {
	Success bool `json:"success"`
}

type GetCodesCouponResponse struct {
	Coupons []CouponResponse `json:"coupons"`
}

type CouponResponse struct {
	ID             string `json:"id"`
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"min_basket_value"`
}

type ApplyCouponResponse struct {
	Basket BasketResponse `json:"basket"`
}

type BasketResponse struct {
	Value                 int  `json:"value"`
	AppliedDiscount       int  `json:"applied_discount"`
	ApplicationSuccessful bool `json:"application_successful"`
}

func MapToGetCodesCouponResponse(coupons []models.Coupon) GetCodesCouponResponse {
	cps := make([]CouponResponse, 0, len(coupons))

	for _, co := range coupons {
		cps = append(cps, CouponResponse{
			ID:             co.ID(),
			Discount:       co.Discount(),
			Code:           co.Code(),
			MinBasketValue: co.MinBasketValue(),
		})
	}

	return GetCodesCouponResponse{Coupons: cps}
}

func MapToApplyCouponResponse(basket models.Basket) ApplyCouponResponse {
	return ApplyCouponResponse{
		Basket: BasketResponse{
			Value:                 basket.Value(),
			AppliedDiscount:       basket.AppliedDiscount(),
			ApplicationSuccessful: basket.ApplicationSuccessful(),
		},
	}
}
