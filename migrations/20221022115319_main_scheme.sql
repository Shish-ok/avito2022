-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_balance (
    user_id BIGINT NOT NULL UNIQUE,
    balance NUMERIC(6, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS balance_holder (
    operation_id BIGINT NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    operation_time TIMESTAMP,
    operation_type TEXT NOT NULL,
    cost NUMERIC(6, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS operation_history (
  operation_id BIGINT NOT NULL UNIQUE,
  user_id BIGINT NOT NULL,
  operation_time TIMESTAMP,
  operation_type TEXT NOT NULL,
  cost NUMERIC(6, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_balance;
DROP TABLE IF EXISTS balance_holder;
DROP TABLE IF EXISTS operation_history;
-- +goose StatementEnd
