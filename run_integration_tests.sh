#!/bin/bash

set -e

# Запускаем docker-compose
echo "Starting docker-compose..."
docker-compose up -d &> /dev/null

# Пауза перед стартом docker-compose
echo "Waiting for 10 seconds before starting docker-compose..."
sleep 10

# Запускаем тесты для всех модулей
echo "Running tests for db-common module..."
# shellcheck disable=SC2164
cd db-common
go test -v ./...

echo "Running tests for api module..."
# shellcheck disable=SC2164
cd ../api
go test -v ./...

echo "Running tests for command-executor module..."
# shellcheck disable=SC2164
cd ../command-executor
go test -v ./...

echo "Running tests for query-handler module..."
# shellcheck disable=SC2164
cd ../query-handler
go test -v ./...

# Останавливаем docker-compose
echo "Stopping docker-compose..."
docker-compose down &> /dev/null

# Удаляем volume
echo "Removing docker volumes..."
docker volume rm signer_db &> /dev/null

echo "All done."