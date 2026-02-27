package repository

import (
	"backend/internal/domain"
	"time"
)



type PaymentRepository interface {
	FindAll() ([]domain.Payment, error)
}

type paymentRepo struct {
	data []domain.Payment
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepo{
		data: []domain.Payment{
			{"p1", "Shopee", time.Now(), 200000, "success"},
			{"p2", "Tokopedia", time.Now(), 150000, "failed"},
			{"p3", "Bukalapak", time.Now(), 500000, "success"},
		},
	}
}

func (r *paymentRepo) FindAll() ([]domain.Payment, error) {
	return r.data, nil
}