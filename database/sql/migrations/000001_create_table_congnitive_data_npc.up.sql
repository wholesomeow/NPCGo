CREATE TABLE IF NOT EXISTS cognitive_data_npc(
    id serial PRIMARY KEY,
    category VARCHAR(128) UNIQUE NOT NULL,
    data_name VARCHAR(256) UNIQUE NOT NULL,
    data_values VARCHAR(128) NOT NULL,
    data_description VARCHAR(256) NOT NULL
);