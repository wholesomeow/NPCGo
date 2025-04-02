CREATE TABLE IF NOT EXISTS jobs_fantasy(
    id serial PRIMARY KEY,
    category VARCHAR(128) UNIQUE NOT NULL,
    job_name VARCHAR(256) UNIQUE NOT NULL,
    alt_name VARCHAR(128),
    job_description VARCHAR(128) NOT NULL,
    can_own BOOL NOT NULL,
    min_status_name VARCHAR(128) NOT NULL,
    max_status_name VARCHAR(128) NOT NULL,
    min_status_level INT NOT NULL,
    max_status_level INT NOT NULL
);