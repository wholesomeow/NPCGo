CREATE TABLE IF NOT EXISTS generator.sex_orientations_npc(
    id serial PRIMARY KEY,
    sexual_orientation VARCHAR(128) UNIQUE NOT NULL,
    sexual_orientation_description VARCHAR(256) UNIQUE NOT NULL
);