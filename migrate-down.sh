#!/bin/sh

wait_for_mysql() {
  until mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" -e "SHOW DATABASES;" >/dev/null 2>&1; do
    echo "Waiting for MySQL..."
    sleep 1
  done
  echo "MySQL started"
}

if [ "$DATABASE" = "mysql" ]; then
  wait_for_mysql
fi

migrate down
migrate -path ./db/migrations -database $DATABASE_URL down
