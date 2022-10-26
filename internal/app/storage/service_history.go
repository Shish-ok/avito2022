package storage

import (
	"avito2022/internal/app/models"
	"context"
	sq "github.com/Masterminds/squirrel"
)

const (
	ServiceTransactionTable = "service_history"
)

func (p *PostgresStorage) AddServiceTransaction(ctx context.Context, transaction models.ServiceTransaction) error {
	query, args, err := sq.
		Insert(ServiceTransactionTable).
		Columns(OrderID, ServiceID, UserID, OperationTime, ServiceName, Cost).
		Values(
			transaction.OrderID,
			transaction.ServiceID,
			transaction.UserID,
			transaction.OperationTime,
			transaction.ServiceName,
			transaction.Cost,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	return err
}
