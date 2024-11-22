#!/bin/bash

if [ -z "$1" ]; then
  echo "Error: Service name (svc) is required."
  exit 1
fi

SVC=$1

SRC_DIR="apps/$SVC"
DEST_DIR="."

if [ ! -d "$SRC_DIR" ]; then
  echo "Error: Service directory $SRC_DIR does not exist."
  exit 1
fi

if [ ! -f "$SRC_DIR/Dockerfile" ]; then
  echo "Error: Dockerfile does not exist in $SRC_DIR."
  exit 1
fi

echo "Moving Dockerfile from $SRC_DIR to $DEST_DIR..."
mv "$SRC_DIR/Dockerfile" "$DEST_DIR/Dockerfile"

echo "Building Docker image for $SVC..."
docker build -t "aiagt-$SVC" "$DEST_DIR"

echo "Moving Dockerfile back to $SRC_DIR..."
mv "$DEST_DIR/Dockerfile" "$SRC_DIR/Dockerfile"

echo "Docker build for $SVC completed."