CREATE TABLE IF NOT EXISTS generator.ngram_fantasy (
    id SERIAL PRIMARY KEY,
    ngram_value VARCHAR(1) UNIQUE NOT NULL,
    ngram_posibility TEXT NOT NULL
);