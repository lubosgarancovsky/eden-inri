#!/bin/bash

# --- Variables for remote access ---
REMOTE_USER=lubos
REMOTE_HOST=pi
REMOTE_PATH=/home/lubos/containers/eden/eden-inri

# Files and folders to sync with remote
FILES_TO_SYNC="build/eden-inri"

# Build Go app
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o build/eden-inri ./cmd/app

# ======================
# Local transfer
# ======================

echo "-- Syncing files with Raspberry Pi"
rsync -avz --delete $FILES_TO_SYNC $REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH/

# Remove build files
rm -rf build

# ======================
# Remote deployment
# ======================
echo "-- Deploying remotely using podman-compose"
ssh $REMOTE_USER@$REMOTE_HOST << EOF
cd $REMOTE_PATH

podman-compose up -d --build --force-recreate

podman image prune -f

EOF