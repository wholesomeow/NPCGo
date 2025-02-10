CREATE TABLE IF NOT EXISTS names_fantasy(
    id serial PRIMARY KEY,
    name_fantasy VARCHAR(128) UNIQUE NOT NULL
);