package storage

import (
	"avito2022/internal/app/models"
	"avito2022/internal/app/service/balance_holder"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
)

const (
	BalanceHolderTable = "balance_holder"
	OrderID            = "order_id"
	ServiceID          = "service_id"
	ServiceName        = "service_name"
)

func (p *PostgresStorage) ReserveMoney(ctx context.Context, operation models.HolderOperation) error {
	tx, err := p.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return err
	}

	defer tx.Rollback()

	balance, err := p.GetUserBalance(ctx, operation.UserID)
	if balance < operation.Cost || err != nil {
		return balance_holder.ErrInsufficientBalance
	}

	if err := p.RemoveMoney(ctx, operation.UserID, operation.Cost); err != nil {
		return err
	}

	query, args, err := sq.
		Insert(BalanceHolderTable).
		Columns(OrderID, ServiceID, UserID, ServiceName, Cost).
		Values(
			operation.OrderID,
			operation.ServiceID,
			operation.UserID,
			operation.ServiceName,
			operation.Cost,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return tx.Commit()
}
