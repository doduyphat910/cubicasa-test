
-- +migrate Up
CREATE TABLE IF NOT EXISTS hubs
(
    id serial NOT NULL,
    name varchar(100) NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS hubs;