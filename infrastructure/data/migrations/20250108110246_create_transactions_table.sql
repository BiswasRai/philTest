-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TYPE payment_status AS ENUM ('PENDING', 'COMPLETED', 'CANCELED');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    payment_status payment_status NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS transactions;
DROP TYPE IF EXISTS payment_status;