package models

type UserBalance struct {
	UserID   uint64  `db:"user_id" json:"user_id"`
	UserName string  `db:"user_name" json:"user_name"`
	Balance  float32 `db:"balance" json:"balance"`
}
