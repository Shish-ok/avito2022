package models

import "time"

type Operation struct {
	OperationID   uint64    `db:"operation_id" json:"operation_id"`
	UserID        uint64    `db:"user_id" json:"user_id"`
	OperationTime time.Time `db:"operation_time" json:"operation_time"`
	OperationType string    `db:"operation_type" json:"operation_type"`
	Cost          float32   `db:"cost" json:"cost"`
}
