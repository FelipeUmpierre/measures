CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS measures CASCADE;
CREATE TABLE events (
    id serial PRIMARY KEY,
    uuid UUID DEFAULT uuid_generate_v1(),
    handle VARCHAR NOT NULL,
    aggregate_id UUID NOT NULL,
    payload JSON,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(id)
);
