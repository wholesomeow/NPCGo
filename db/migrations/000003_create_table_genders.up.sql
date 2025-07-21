CREATE TABLE IF NOT EXISTS genders(
    id serial PRIMARY KEY,
    gender VARCHAR(128) NOT NULL,
    gender_description VARCHAR(256) NOT NULL,
    pronouns VARCHAR(128) NOT NULL,
    secondary_pronouns VARCHAR(128),
    tirtiary_pronouns VARCHAR(128)
);