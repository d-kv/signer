#!/bin/bash

# Запускаем docker-compose
echo "Starting docker-compose..."
docker-compose up -d &> /dev/null

# Пауза перед стартом docker-compose
echo "Waiting for 5 seconds before starting docker-compose..."
sleep 5

# Меняем директорию и запускаем тесты
echo "Running tests..."
# shellcheck disable=SC2164
cd repo/domain/
go test -v

# Останавливаем docker-compose
echo "Stopping docker-compose..."
docker-compose down &> /dev/null

# Удаляем volume
echo "Removing docker volumes..."
docker volume rm signer_db &> /dev/null

echo "All done."