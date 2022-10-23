package balance

import (
	"avito2022/internal/app/models"
	"context"
	"errors"
)

var (
	ErrInvalidBalance = errors.New("invalid balance")
)

type Storage interface {
	UpsetUserBalance(context.Context, models.UserBalance) error
	GetUserBalance(context.Context, uint64) (float32, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) UpsetUserBalance(ctx context.Context, userBalance models.UserBalance) error {
	if userBalance.Balance < 0 {
		return ErrInvalidBalance
	}
	return s.storage.UpsetUserBalance(ctx, userBalance)
}

func (s *Service) GetUserBalance(ctx context.Context, userID uint64) (float32, error) {
	return s.storage.GetUserBalance(ctx, userID)
}
