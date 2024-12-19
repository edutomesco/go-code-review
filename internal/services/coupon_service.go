package services

import (
	"github.com/edutomesco/coupons/internal/interfaces"
	"github.com/edutomesco/coupons/internal/models"
	"github.com/edutomesco/coupons/internal/services/dto"

	"github.com/google/uuid"
)

type CouponService struct {
	repo interfaces.CouponRepository
}

func NewCouponService(repo interfaces.CouponRepository) *CouponService {
	return &CouponService{
		repo: repo,
	}
}

func (s *CouponService) CreateCoupon(req dto.CreateCouponRequest) error {
	coupon, err := models.NewCoupon(uuid.NewString(), req.Discount, req.Code, req.MinBasketValue)
	if err != nil {
		return err
	}

	return s.repo.Save(coupon)
}

func (s *CouponService) GetCoupons(codes []string) ([]models.Coupon, error) {
	coupons := make([]models.Coupon, 0, len(codes))

	for _, code := range codes {
		coupon, err := s.repo.GetByCode(code)
		if err != nil {
			return nil, err
		}
		coupons = append(coupons, coupon)
	}

	return coupons, nil
}

func (s *CouponService) ApplyCoupon(req dto.ApplyCouponRequest) (models.Basket, error) {
	coupon, err := s.repo.GetByCode(req.Code)
	if err != nil {
		return models.Basket{}, err
	}

	return models.NewBasket(req.Basket.Value, coupon)
}
