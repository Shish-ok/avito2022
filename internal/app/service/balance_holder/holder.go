package balance_holder

import (
	"avito2022/internal/app/models"
	"context"
	"errors"
)

var (
	ErrInvalidCost         = errors.New("invalid cost")
	ErrInsufficientBalance = errors.New("insufficient balance")
)

type Storage interface {
	ReserveMoney(context.Context, models.HolderOperation) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) ReserveMoney(ctx context.Context, operation models.HolderOperation) error {
	if operation.Cost < 0 {
		return ErrInvalidCost
	}

	return s.storage.ReserveMoney(ctx, operation)
}
