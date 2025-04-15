#!/bin/bash

if [ -z $1 ]; then
echo "Must provide version to force DB back to"
exit 1
fi

CONFIG_FILE="./configs/dbconf.yml"

# Read in server values from the YAML file
SERVER_PORT=$(yq '.server.port' $CONFIG_FILE)
SERVER_LOGLEVEL=$(yq '.server.loglevel' $CONFIG_FILE)
SERVER_NETWORK=$(yq '.server.network' $CONFIG_FILE)

# Read in database values from the YAML file
DB_NAME=$(yq '.database.dbname' $CONFIG_FILE | sed 's|"||g')
DB_USER=$(yq '.database.user' $CONFIG_FILE | sed 's|"||g')
DB_PW=$(yq '.database.password' $CONFIG_FILE | sed 's|"||g')
DB_PORT=$(yq '.database.port' $CONFIG_FILE)
DB_HOST=$(yq '.database.hostname' $CONFIG_FILE)

echo "Variables read from config file. Writing to .env"
CONTAINER_NAMES=(
  npcg-pg
  npcg-app
  npcg-pg-migrate
  )

# Setting MIGRATION_PATH here because I don't want to try and split "file://" from the path in the dbconf.yml
MIGRATION_PATH=./internal/sql/migrations
RAWDATA_PATH=./data/raw

echo "Setting DB connection string"
DB_CONNECTION_STRING=postgres://${DB_USER}:${DB_PW}@${CONTAINER_NAMES[0]}:${DB_PORT}/${DB_NAME}?sslmode=disable

# Force DB back to specific version
echo "--- Attempting to FORCE database version back to "$1" ---"
echo "Starting ${CONTAINER_NAMES[0]}"
docker network create npcg-network > /dev/null
docker run -d \
--name ${CONTAINER_NAMES[0]} \
--restart always \
--env-file ./.env \
-e POSTGRES_USER=${DB_USER} \
-e POSTGRES_DB=${DB_NAME} \
-e POSTGRES_PASSWORD=${DB_PW} \
-p 5432:5432 \
-v ./pg-data:/var/lib/postgresql/data \
--health-cmd="pg_isready -U ${DB_USER} -d ${DB_NAME}" \
--health-interval=10s \
--health-timeout=5s \
--health-retries=5 \
--network npcg-network \
postgres:10.5 > /dev/null || exit 1

echo "Current database version:"
docker run --rm \
--network npcg-network \
-v ${MIGRATION_PATH}:/migrations \
migrate/migrate \
-path=/migrations \
-database ${DB_CONNECTION_STRING} \
version || exit 1

echo "Attempting roll back"
docker run --rm \
--network npcg-network \
-v ${MIGRATION_PATH}:/migrations \
migrate/migrate \
-path=/migrations \
-database ${DB_CONNECTION_STRING} \
force $3 || exit 1

echo "Stopping ${CONTAINER_NAMES[0]}"
docker stop ${CONTAINER_NAMES[0]}
docker rm ${CONTAINER_NAMES[0]}
docker network remove npcg-network
echo "--- Rollback Attempt Completed ---"