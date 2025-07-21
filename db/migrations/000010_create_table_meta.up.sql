CREATE TABLE npc_types (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    enum TEXT
);

CREATE TABLE body_types (
    id UUID PRIMARY KEY,
    name TEXT,
    enum TEXT
);

CREATE TABLE sex_types (
    id UUID PRIMARY KEY,
    name TEXT,
    enum TEXT
);

CREATE TABLE gender_types (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    enum TEXT
);

CREATE TABLE orientation_types (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    enum TEXT
);