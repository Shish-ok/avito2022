-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXIST user_balance (
    user_id BIGINT NOT NULL UNIQUE,
    user_name TEXT NOT NULL,
    balance DOUBLE NOT NULL
);

CREATE TABLE IF NOT EXIST balance_holder (
    operation_id BIGINT NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    operation_time TIMESTAMP,
    operation_type TEXT NOT NULL,
    cost DOUBLE NOT NULL
);

CREATE TABLE IF NOT EXIST operation_history (
  operation_id BIGINT NOT NULL UNIQUE,
  user_id BIGINT NOT NULL,
  operation_time TIMESTAMP,
  operation_type TEXT NOT NULL,
  cost DOUBLE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXIST user_balance;
DROP TABLE IF EXIST balance_holder;
DROP TABLE IF EXIST operation_history;
-- +goose StatementEnd
