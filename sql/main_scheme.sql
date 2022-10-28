CREATE TABLE user_balance (
    user_id BIGINT NOT NULL UNIQUE,
    balance NUMERIC(6, 2) NOT NULL
);

CREATE TABLE balance_holder (
    order_id BIGINT NOT NULL UNIQUE,
    service_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    service_name TEXT NOT NULL,
    cost NUMERIC(6, 2) NOT NULL
);

CREATE TABLE service_history (
    order_id BIGINT NOT NULL UNIQUE,
    service_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    operation_time DATE,
    service_name TEXT NOT NULL,
    cost NUMERIC(6, 2) NOT NULL
);

CREATE TABLE transaction_history (
    transaction_id BIGINT NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    operation_time TIMESTAMP,
    description TEXT NOT NULL,
    cost NUMERIC(6, 2) NOT NULL
);

CREATE UNIQUE INDEX pagination_time
    ON transaction_history (operation_time, transaction_id);

CREATE UNIQUE INDEX pagination_cost
    ON transaction_history (cost, transaction_id);