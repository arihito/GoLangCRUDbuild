#!/bin/bash

# exit immidiately when a command fails
set -e

if [ -z "$MYSQL_DB_HOST" ]; then
    >&2 echo "db host undefined"
    exit 1
fi
db_host="$MYSQL_DB_HOST"

if [ -z "$MYSQL_USER" ]; then
    >&2 echo "mysql user undefined"
    exit 1
fi
db_user="$MYSQL_USER"

if [ -z "$MYSQL_PASSWORD" ]; then
    >&2 echo "mysql password undefined"
    exit 1
fi
db_password="$MYSQL_PASSWORD"

until mysqladmin ping -u "$db_user" -h "$db_host" -p"$db_password"; do
  >&2 echo "mysql is unavailable - sleeping"
  sleep 1
done

>&2 echo "mysql is up - executing command"

# start...
bash -c "realize start --name='sample-api' --run"
