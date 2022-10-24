package storage

import (
	"avito2022/internal/app/models"
	"context"
	sq "github.com/Masterminds/squirrel"
)

const (
	UserTransactionsTable = "transaction_history"
	TransactionID         = "transaction_id"
	OperationTime         = "operation_time"
	Description           = "description"
	Cost                  = "cost"
)

func (p *PostgresStorage) MakeUserTransaction(ctx context.Context, transaction models.UserTransaction) error {
	query, args, err := sq.
		Insert(UserTransactionsTable).
		Columns(TransactionID, UserID, OperationTime, Description, Cost).
		Values(
			transaction.TransactionID,
			transaction.UserID,
			transaction.OperationTime,
			transaction.Description,
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
