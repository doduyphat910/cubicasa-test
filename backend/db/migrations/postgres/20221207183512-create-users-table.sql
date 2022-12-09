
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id serial NOT NULL,
    team_id integer,
    CONSTRAINT fk_users_team_id FOREIGN KEY (team_id) REFERENCES teams (id),
    type varchar(100) NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    PRIMARY KEY (id)
    );

-- +migrate Down
DROP TABLE IF EXISTS users;