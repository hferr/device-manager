-- +goose Up
CREATE TYPE device_states AS ENUM('available', 'in_use', 'inactive');

CREATE TABLE devices(
    id uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    state device_states NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TYPE IF EXISTS device_states;
DROP TABLE IF EXISTS devices;
