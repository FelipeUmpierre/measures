CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR,
    UNIQUE(id)
);

CREATE TABLE measures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    measures JSON,
    UNIQUE(id)
);
