package services

import (
	"errors"
	mock_interfaces "github.com/edutomesco/coupons/internal/interfaces/mocks"
	"github.com/edutomesco/coupons/internal/models"
	"github.com/edutomesco/coupons/internal/services/dto"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_CreateCoupon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_interfaces.NewMockCouponRepository(ctrl)
	cs := NewCouponService(cr)

	couponRequest := dto.CreateCouponRequest{
		Code:           "xxx",
		Discount:       30,
		MinBasketValue: 60,
	}

	t.Run("Given no database error, when saving a coupon, should work", func(t *testing.T) {
		cr.EXPECT().Save(gomock.Any()).Return(nil)

		err := cs.CreateCoupon(couponRequest)
		assert.NoError(t, err)
	})

	t.Run("Given an invalid coupon format, when saving a coupon, should return error", func(t *testing.T) {
		invalidCouponRequest := dto.CreateCouponRequest{
			Code:           "xxx",
			Discount:       30,
			MinBasketValue: 29,
		}

		err := cs.CreateCoupon(invalidCouponRequest)
		assert.Error(t, err)
		assert.Equal(t, models.ErrInvalidMinBasketValue, err)
	})

	t.Run("Given a database error, when saving a coupon, should return error", func(t *testing.T) {
		cr.EXPECT().Save(gomock.Any()).Return(errors.New("failed repository"))

		err := cs.CreateCoupon(couponRequest)
		assert.Error(t, err)
	})
}

func TestService_GetCoupons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_interfaces.NewMockCouponRepository(ctrl)
	cs := NewCouponService(cr)

	code1 := "xxx"
	code2 := "xxxx"
	codes := []string{code1, code2}

	coupon1, err := models.NewCoupon(uuid.NewString(), 10, code1, 300)
	require.NoError(t, err)
	coupon2, err := models.NewCoupon(uuid.NewString(), 20, code2, 600)
	require.NoError(t, err)
	expectedCoupons := []models.Coupon{coupon1, coupon2}

	t.Run("Given existing coupon codes, when fetching coupons, should work", func(t *testing.T) {
		cr.EXPECT().GetByCode(gomock.Any()).Return(coupon1, nil)
		cr.EXPECT().GetByCode(gomock.Any()).Return(coupon2, nil)

		coupons, err := cs.GetCoupons(codes)
		assert.NoError(t, err)
		assert.ElementsMatch(t, expectedCoupons, coupons)
	})

	t.Run("Given no existing coupon code, when fetching coupons, should return error", func(t *testing.T) {
		cr.EXPECT().GetByCode(gomock.Any()).Return(models.Coupon{}, errors.New("not found"))

		_, err = cs.GetCoupons(codes)
		assert.Error(t, err)
	})
}

func TestService_ApplyCoupon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_interfaces.NewMockCouponRepository(ctrl)
	cs := NewCouponService(cr)

	code := "xxx"
	basketValue := 100
	minBasketValue := 51
	req := dto.ApplyCouponRequest{
		Code:   code,
		Basket: dto.BasketRequest{Value: basketValue},
	}
	coupon, err := models.NewCoupon(uuid.NewString(), 50, code, minBasketValue)
	require.NoError(t, err)

	t.Run("Given correct coupon basket, when applying coupon, should work", func(t *testing.T) {
		expectedBasketValue := 50

		cr.EXPECT().GetByCode(gomock.Any()).Return(coupon, nil)

		basket, err := cs.ApplyCoupon(req)
		assert.NoError(t, err)
		assert.Equal(t, expectedBasketValue, basket.Value())
		assert.Equal(t, true, basket.ApplicationSuccessful())
		assert.Equal(t, coupon.Discount(), basket.AppliedDiscount())
	})

	t.Run("Given invalid basket value, when applying coupon, should return error", func(t *testing.T) {
		invalidReq := dto.ApplyCouponRequest{
			Code:   code,
			Basket: dto.BasketRequest{Value: 50},
		}

		cr.EXPECT().GetByCode(gomock.Any()).Return(coupon, nil)

		_, err = cs.ApplyCoupon(invalidReq)
		assert.Error(t, err)
		assert.Equal(t, models.ErrInsufficientBasketValue, err)
	})

	t.Run("Given no existent code coupon, when applying coupon, should return error", func(t *testing.T) {
		cr.EXPECT().GetByCode(gomock.Any()).Return(models.Coupon{}, errors.New("not found"))

		_, err = cs.ApplyCoupon(req)
		assert.Error(t, err)
	})
}
