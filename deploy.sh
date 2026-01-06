#!/bin/bash

# --- Variables for remote access ---
REMOTE_USER=lubos
REMOTE_HOST=pi
REMOTE_PATH=/home/lubos/containers/eden/eden-inri
SERVICE_NAME='eden-inri'

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

echo "-- Stop and disable service"
systemctl --user stop container-$SERVICE_NAME.service
systemctl --user disable container-$SERVICE_NAME.service

echo "-- Building the service image"
podman-compose up -d --build --force-recreate

echo "-- Generating systemd unit"
podman generate systemd --name $SERVICE_NAME --files --new
mkdir -p ~/.config/systemd/user/
mv container-$SERVICE_NAME.service ~/.config/systemd/user/
podman stop $SERVICE_NAME

echo "-- Staring and enabling service"
systemctl --user daemon-reload
systemctl --user enable container-$SERVICE_NAME.service
systemctl --user start container-$SERVICE_NAME.service

podman image prune -f

EOF