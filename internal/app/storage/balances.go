package storage

import (
	"avito2022/internal/app/models"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

const (
	BalanceTable = "user_balance"
	UserID       = "user_id"
	Balance      = "balance"
)

func (p *PostgresStorage) UpsetUserBalance(ctx context.Context, balance models.UserBalance) error {
	onConflict := fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s = %s.%s + $2", UserID, Balance, BalanceTable, Balance)
	query, args, err := sq.
		Insert(BalanceTable).
		Columns(UserID, Balance).
		Values(
			balance.UserID,
			balance.Balance,
		).
		Suffix(onConflict).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	return err

}

func (p *PostgresStorage) GetUserBalance(ctx context.Context, userID uint64) (float32, error) {
	var balance float32
	query, args, err := sq.
		Select(Balance).
		From(BalanceTable).
		Where(sq.Eq{UserID: userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return balance, err
	}

	err = p.db.GetContext(ctx, &balance, query, args...)
	return balance, err
}
