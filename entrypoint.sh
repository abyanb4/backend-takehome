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

# On first build, usually MySQL doesn't ready to accept connection immediately
# Hence, if the migration fail, wait 3 second  until the database hopefully ready
run_migration() {
  local attempt=1
  local max_attempts=2

  while [ $attempt -le $max_attempts ]; do
    echo "Attempt $attempt: Running migration..."
    migrate -path /app/db/migrations -database $DB_URL up

    if [ $? -eq 0 ]; then
      echo "Migration succeeded."
      return 0
    else
      echo "Migration failed. Retrying in 3 seconds..."
      sleep 3
    fi

    attempt=$((attempt + 1))
  done

  echo "Migration failed after $max_attempts attempts."
  return 1
}

run_migration

exec /app/main
