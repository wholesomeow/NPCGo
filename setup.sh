#!/bin/bash

if [ -z $1 ]; then
  echo "Must provide mode as argument"
  echo "Options are: dev or prod"
  exit 1
fi

CONFIG_FILE="./configuration/dbconf.yml"

# Read in server values from the YAML file
SERVER_PORT=$(yq '.server.port' $CONFIG_FILE)
SERVER_LOGLEVEL=$(yq '.server.loglevel' $CONFIG_FILE)
SERVER_NETWORK=$(yq '.server.network' $CONFIG_FILE)

# Read in database values from the YAML file
DB_NAME=$(yq '.database.dbname' $CONFIG_FILE)
DB_USER=$(yq '.database.user' $CONFIG_FILE)
DB_PW=$(yq '.database.password' $CONFIG_FILE)
DB_PORT=$(yq '.database.port' $CONFIG_FILE)
DB_HOST=$(yq '.database.hostname' $CONFIG_FILE)

# Determine mode
case $1 in
  dev )
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

echo "Variables read from config file. Writing to .env"
CONTAINER_NAMES=(
  npcg-pg
  npcg-app
  npcg-pg-migrate
  )

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
EOF

# echo "Cleaning containers"
# docker compose down --remove-orphans

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

# echo "Cleaning workspace"
# sudo chmod -r $USER:$USER postgres-data

# Will probably end up commenting this part out
# sudo rm -rf postgres-data

echo "Starting new build"
docker compose up --build