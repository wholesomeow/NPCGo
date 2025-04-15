CREATE TABLE IF NOT EXISTS names_fantasy(
    id serial PRIMARY KEY,
    letter VARCHAR(1) NOT NULL,
    name_value VARCHAR(128) UNIQUE NOT NULL
);