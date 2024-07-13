#!/bin/bash
set -e

if [[ "$COMPOSE" == "true" ]]; then
  cd docker-entrypoint-initdb.d
fi

psql -U "$POSTGRES_USER" -f sql/00-start-scripts.sql "$POSTGRES_DB"