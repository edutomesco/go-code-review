package memdb

import (
	"errors"
	"fmt"
	"github.com/edutomesco/coupons/internal/models"
	appError "github.com/edutomesco/coupons/internal/models/errors"
	"sync"
)

type CouponMemoryEntity struct {
	ID             string
	Discount       int
	Code           string
	MinBasketValue int
}

func mapToCouponMemoryEntity(co models.Coupon) CouponMemoryEntity {
	return CouponMemoryEntity{
		ID:             co.ID(),
		Discount:       co.Discount(),
		Code:           co.Code(),
		MinBasketValue: co.MinBasketValue(),
	}
}

func mapToCoupon(co CouponMemoryEntity) (models.Coupon, error) {
	return models.NewCoupon(co.ID, co.Discount, co.Code, co.MinBasketValue)
}

type CouponRepository struct {
	mu      sync.Mutex
	entries map[string]CouponMemoryEntity
}

func NewCouponRepository() *CouponRepository {
	return &CouponRepository{
		entries: make(map[string]CouponMemoryEntity, 0),
	}
}

func (r *CouponRepository) Save(coupon models.Coupon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.entries[coupon.Code()]
	if ok {
		return appError.ErrUnexpected(errors.New("coupon already exists"))
	}

	memCoupon := mapToCouponMemoryEntity(coupon)

	r.entries[coupon.Code()] = memCoupon

	return nil
}

func (r *CouponRepository) GetByCode(code string) (models.Coupon, error) {
	memCoupon, ok := r.entries[code]
	if !ok {
		return models.Coupon{}, appError.ErrComponentNotFound(fmt.Sprintf("code %s", code))
	}

	return mapToCoupon(memCoupon)
}
