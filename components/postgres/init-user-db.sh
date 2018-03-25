#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER felipe;
    CREATE DATABASE measures;
    GRANT ALL PRIVILEGES ON DATABASE measures TO felipe;
EOSQL