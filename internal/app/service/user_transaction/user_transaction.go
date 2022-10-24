package user_transaction

import (
	"avito2022/internal/app/models"
	"context"
	"log"
)

const (
	UpBalance = "Пополнение баланса"
)

type Storage interface {
	MakeUserTransaction(context.Context, models.UserTransaction) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) UpBalanceTransaction(ctx context.Context, userID uint64, cost float32) {
	err := s.storage.MakeUserTransaction(ctx, models.NewUserTransaction(userID, UpBalance, cost))
	if err != nil {
		log.Print(err)
	}
}
