#!/bin/bash

# this file needs to be executed with sudo

cat <<EOF | tee /etc/systemd/system/docker-compose-gpc.service > /dev/null
[Unit]
Description=Docker Compose Application Service
Requires=docker.service
After=docker.service

[Service]
Type=oneshot
RemainAfterExit=yes
WorkingDirectory=/app/github-profile-card
Environment=PATH=/usr/bin:/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin
ExecStart=/usr/bin/make up
ExecStop=/usr/bin/make down
TimeoutStartSec=0

[Install]
WantedBy=multi-user.target
EOF

systemctl enable docker-compose-gpc
systemctl start docker-compose-gpc
systemctl status docker-compose-gpc

