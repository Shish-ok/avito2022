package bugalter_accounting

import (
	"avito2022/internal/app/models"
	"context"
	"log"
)

type Storage interface {
	AddServiceTransaction(context.Context, models.ServiceTransaction) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) AddServiceTransaction(ctx context.Context, operation models.HolderOperation) {
	err := s.storage.AddServiceTransaction(ctx, models.NewServiceTransaction(operation))
	if err != nil {
		log.Print(err)
	}
}
