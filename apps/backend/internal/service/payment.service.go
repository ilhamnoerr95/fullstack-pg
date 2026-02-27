package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
)

type PaymentService interface {
	GetAll() ([]domain.Payment, map[string]int, error)
	GetPaymentsByStatus(status string) ([]domain.Payment, error)
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

func (s *paymentService) GetPaymentsByStatus(status string) ([]domain.Payment, error) {

	switch domain.PaymentStatus(status) {
	case domain.StatusCompleted,
		domain.StatusProcessing,
		domain.StatusFailed:
	default:
		return nil, errors.New("invalid payment status")
	}

	return s.repo.FindByStatus(domain.PaymentStatus(status))
}