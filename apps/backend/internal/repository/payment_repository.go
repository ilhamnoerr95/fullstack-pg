package repository

import (
	"backend/internal/domain"
	"time"
)



type PaymentRepository interface {
	FindAll() ([]domain.Payment, error)
	FindByStatus(status domain.PaymentStatus) ([]domain.Payment, error)
}

type paymentRepo struct {
	data []domain.Payment
}

// static
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

func (r *paymentRepo) FindByStatus(status domain.PaymentStatus) ([]domain.Payment, error) {
	var result []domain.Payment
	for _, p := range r.data {
		if p.Status == status {
			result = append(result, p)
		}
	}
	return result, nil
}