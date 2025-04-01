CREATE TABLE IF NOT EXISTS enneagram(
    id serial PRIMARY KEY,
    archetype VARCHAR(128) UNIQUE NOT NULL,
    center VARCHAR(256) UNIQUE NOT NULL,
    dominant_emtion VARCHAR(128) NOT NULL,
    keywords VARCHAR(128) NOT NULL,
    enneagram_description VARCHAR(256) NOT NULL,
    fear VARCHAR(128),
    desire VARCHAR(128) NOT NULL,
    wings VARCHAR(128) NOT NULL,
    lod_level VARCHAR(128) NOT NULL,
    current_lod VARCHAR(128) NOT NULL,
    level_of_development VARCHAR(128) NOT NULL,
    key_motivations VARCHAR(128) NOT NULL,
    overview VARCHAR(128) NOT NULL,
    addictions VARCHAR(128) NOT NULL,
    growth_recommendations VARCHAR(128) NOT NULL
);