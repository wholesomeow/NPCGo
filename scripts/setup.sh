#!/bin/bash

# Check arguments
if [ -z $1 ]; then
  echo "Must provide mode as argument"
  echo "Options are: dev, build, or prod"
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

# Determine mode
case $1 in
  dev )
    ENV="development"
    ;;
  build )
    ENV="development"
    ;;
  prod )
    ENV="production"
    ;;
  * )
    echo "Unknown environment: $1"
    exit 1
    ;;
esac

# Write config file variables to .env file for docker-compose
cat > .env << EOF
GO_ENV=${ENV}
NPCG_PORT=${SERVER_PORT}
DB_PORT=${DB_PORT}
LOG_LEVEL=${SERVER_LOGLEVEL}
NETWORK=${SERVER_NETWORK}
POSTGRES_DB=${DB_NAME}
POSTGRES_USER=${DB_USER}
POSTGRES_PASSWORD=${DB_PW}
POSTGRES_HOST=${DB_HOST}
POSTGRES_CONTAINER_NAME=${CONTAINER_NAMES[0]}
APP_CONTAINER_NAME=${CONTAINER_NAMES[1]}
MIGRATION_CONTAINER_NAME=${CONTAINER_NAMES[2]}
MIGRATION_PATH=${MIGRATION_PATH}
RAWDATA_PATH=${RAWDATA_PATH}
DB_CONNECTION_STRING=${DB_CONNECTION_STRING}
EOF

echo "Current ENV file values"
echo  "GO_ENV: ${ENV}"
echo  "NPCG_PORT: ${SERVER_PORT}"
echo  "DB_PORT: ${DB_PORT}"
echo  "NETWORK: ${SERVER_NETWORK}"
echo  "POSTGRES_DB: ${DB_NAME}"
echo  "POSTGRES_USER: ${DB_USER}"
echo  "POSTGRES_HOST: ${DB_HOST}"
echo  "POSTGRES_CONTAINER_NAME: ${CONTAINER_NAMES[0]}"
echo  "APP_CONTAINER_NAME: ${CONTAINER_NAMES[1]}"
echo  "MIGRATION_CONTAINER_NAME: ${CONTAINER_NAMES[2]}"
echo  "MIGRATION_PATH: ${MIGRATION_PATH}"
echo  "RAWDATA_PATH: ${RAWDATA_PATH}"
echo  "DB_CONNECTION_STRING: ${DB_CONNECTION_STRING}"

echo "Checking space in /var"
VAR_OUTPUT=$(sudo du -cha --max-depth=1 /var | grep -E "M|G" | tail -n 1)
VAR_SIZE=$(echo "$VAR_OUTPUT" | awk '{print $1}' | sed 's/[[:alpha:]]//g')
VAR_SIZE_FLOAT=$(echo "$VAR_SIZE" | awk '{printf "%.2f", $1}')
VAR_THREASHOLD=12.0

echo "Directory /var current size: $VAR_SIZE_FLOAT"
if (( $(echo "$VAR_SIZE_FLOAT > $VAR_THREASHOLD" | bc -l) )); then
  echo "Directory /var getting too large... running docker prune"
  docker system prune -a -f
fi

echo "Checking for existing running containers"
for NAME in ${CONTAINER_NAMES[@]}; do
    if docker ps -a --filter "name=$NAME" --format "{{.Names}}" | grep -w "$NAME" > /dev/null; then
        echo "Container '$NAME' exists. Removing it..."
        docker stop "$NAME" > /dev/null
        docker rm "$NAME" > /dev/null
    else
        echo "Container '$NAME' does not exist."
    fi
done

# Determine docker compose command
echo "Starting up containers"
case $1 in
  dev )
    docker compose up --no-recreate
    ;;
  build )
    docker compose up --build
    ;;
  prod )
    docker compose up --build
    ;;
  * )
    echo "Unknown environment: $1 failed to trigger docker compose command"
    exit 1
    ;;
esac