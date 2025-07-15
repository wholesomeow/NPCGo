CREATE TABLE IF NOT EXISTS users.users (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    email_verified BOOLEAN DEFAULT FALSE,
    verification_token TEXT,
    verification_sent_at TIMESTAMP,
    verified_at TIMESTAMP,

    tier TEXT NOT NULL DEFAULT 'free',  -- free, standard, pro, commercial, devkit
    acct_type TEXT NOT NULL DEFAULT 'user', -- user, admin
    acct_status TEXT NOT NULL DEFAULT 'active',  -- active, suspended, deleted

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    last_login_at TIMESTAMP,
    metadata JSONB DEFAULT '{}',

    -- Indexes for performance
    UNIQUE (email),
    UNIQUE (username)
);