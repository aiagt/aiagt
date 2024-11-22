#!/bin/bash

services=("gateway" "user" "plugin" "app" "chat" "model")

current_dir=$(basename "$PWD")
if [ "$current_dir" == "deploy" ]; then
    cd ..
fi

for svc in "${services[@]}"; do
  "./deploy/docker-build.sh" "$svc"
  echo "Build $svc success"
done