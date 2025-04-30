CREATE TABLE npc_meta.npc_types (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    enum TEXT
);

CREATE TABLE npc_meta.body_types (
    id UUID PRIMARY KEY,
    name TEXT,
    enum TEXT
);

CREATE TABLE npc_meta.sex_types (
    id UUID PRIMARY KEY,
    name TEXT,
    enum TEXT
);

CREATE TABLE npc_meta.gender_types (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    enum TEXT
);

CREATE TABLE npc_meta.orientation_types (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    enum TEXT
);