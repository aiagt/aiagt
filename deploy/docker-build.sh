#!/bin/bash

if [ -z "$1" ]; then
  echo "Error: Service name (svc) is required."
  exit 1
fi

SVC=$1

DOCKER_FILE="apps/$SVC/Dockerfile"

if [ ! -f "$DOCKER_FILE" ]; then
  echo "Error: Dockerfile does not exist."
  exit 1
fi

echo "Building Docker image for $SVC..."
docker build -t "aiagt-$SVC" -f "$DOCKER_FILE" .

echo "Docker build for $SVC completed."