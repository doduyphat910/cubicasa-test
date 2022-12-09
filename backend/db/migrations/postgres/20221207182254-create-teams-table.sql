
-- +migrate Up
CREATE TABLE IF NOT EXISTS teams
(
    id serial NOT NULL,
    hub_id integer NOT NULL,
    CONSTRAINT fk_teams_hub_id FOREIGN KEY (hub_id) REFERENCES hubs (id),
    geo_location point NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS teams;
