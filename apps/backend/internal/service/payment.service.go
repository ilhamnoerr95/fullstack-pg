package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"backend/internal/domain"
	"backend/internal/lib"
	"backend/internal/repository"

	"github.com/redis/go-redis/v9"
)

// business prccess
type PaymentService struct {
	repo  repository.PaymentRepository
	redis *redis.Client
}

func NewPaymentService(repo repository.PaymentRepository, redis *redis.Client) *PaymentService {
	return &PaymentService{
		repo:  repo,
		redis: redis,
	}
}

func (s *PaymentService) GetAll() ([]domain.Payment, map[string]int, error) {

	ctx := context.Background()
	cacheKey := "payments:all"

	// 1 Try Redis
	val, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var cached []domain.Payment
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			summary := lib.CalculateSummary(cached)
			return cached, summary, nil
		}
	}

	// 2️ Fallback DB
	payments, err := s.repo.FindAll()
	if err != nil {
		return nil, nil, err
	}

	// 3️ Calculate summary
	summary := lib.CalculateSummary(payments)

	// 4️ Save to Redis
	bytes, _ := json.Marshal(payments)
	_ = s.redis.Set(ctx, cacheKey, bytes, 5*time.Minute).Err()

	return payments, summary, nil
}

func (s *PaymentService) GetPaymentsByStatus(status string) ([]domain.Payment, error) {

	switch domain.PaymentStatus(status) {
	case domain.StatusCompleted,
		domain.StatusProcessing,
		domain.StatusFailed:
	default:
		return nil, errors.New("invalid payment status")
	}

	return s.repo.FindByStatus(domain.PaymentStatus(status))
}