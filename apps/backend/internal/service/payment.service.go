package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type PaymentService interface {
	GetAll() ([]domain.Payment, map[string]int, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(r repository.PaymentRepository) PaymentService {
	return &paymentService{r}
}

func (s *paymentService) GetAll() ([]domain.Payment, map[string]int, error) {
	data, _ := s.repo.FindAll()

	summary := map[string]int{
		"total":   len(data),
		"success": 0,
		"failed":  0,
	}

	for _, p := range data {
		if p.Status == "success" {
			summary["success"]++
		}
		if p.Status == "failed" {
			summary["failed"]++
		}
	}

	return data, summary, nil
}