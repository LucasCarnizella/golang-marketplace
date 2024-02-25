#!/bin/bash
set -e

if [[ $COMPOSE == true ]]
then
  cd docker-entrypoint-initdb.d || exit 1
fi

psql -U $DB_USER -f sql/00-start-scripts.sql $DB_NAME