#!/bin/bash

# --- Variables for remote access ---
REMOTE_USER=lubos
REMOTE_HOST=pi
REMOTE_PATH=/home/lubos/homelab/apps/eden/eden-inri
SYSTEMD_PATH=/home/lubos/.config/systemd/user
SERVICE_NAME='eden-inri'

# Files and folders to sync with remote
FILES_TO_SYNC="build/eden-inri Dockerfile bin/eden-inri.service"

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
systemctl --user stop $SERVICE_NAME.service
systemctl --user disable $SERVICE_NAME.service

echo "-- Creating systemd service"
mv $SERVICE_NAME.service $SYSTEMD_PATH

echo "-- Staring and enabling service"
systemctl --user daemon-reload
systemctl --user enable $SERVICE_NAME.service
systemctl --user start $SERVICE_NAME.service

podman image prune -f

EOF