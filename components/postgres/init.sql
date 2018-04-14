CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    name VARCHAR,
    UNIQUE(id)
);

DROP TABLE IF EXISTS events CASCADE;
CREATE TABLE events (
    id serial PRIMARY KEY,
    handle VARCHAR NOT NULL,
    version serial NOT NULL,
    aggregate_id UUID NOT NULL,
    payload JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(id)
);

INSERT INTO users (name) VALUES ('Felipe'), ('Vivian');
