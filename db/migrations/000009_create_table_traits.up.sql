CREATE TABLE ocean_data (
    id UUID PRIMARY KEY,
    aspect TEXT[],
    degree FLOAT8[],
    traits TEXT[][],
    text TEXT,
    description TEXT[],
    use TEXT
);

CREATE TABLE mice_data (
    id UUID PRIMARY KEY,
    aspect TEXT,
    degree FLOAT8,
    traits TEXT[],
    text TEXT,
    description TEXT,
    use TEXT
);

CREATE TABLE rei_data (
    id UUID PRIMARY KEY,
    aspect TEXT[],
    degree FLOAT8[],
    traits TEXT[],
    text TEXT,
    description TEXT[],
    use TEXT
);

CREATE TABLE cs_data (
    id UUID PRIMARY KEY,
    aspect TEXT,
    traits TEXT[],
    text TEXT,
    coords INTEGER[2],
    description TEXT,
    use TEXT
);

CREATE TABLE enneagram_data (
    id UUID PRIMARY KEY,
    enneagram_id INT,
    archetype TEXT,
    keywords TEXT[],
    description TEXT,
    center TEXT,
    dominant_emotion TEXT,
    fear TEXT,
    desire TEXT,
    wings INTEGER[2],
    lod_level INT,
    current_lod TEXT,
    level_of_development TEXT[9],
    key_motivations TEXT,
    overview TEXT,
    addictions TEXT,
    growth_recommendations TEXT[5]
);