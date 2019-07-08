#!/bin/bash

# Get all project dependencies
go get gotest.tools
go get github.com/google/go-cmp
go get github.com/pkg/errors
go get github.com/mattn/go-sqlite3
go get github.com/Murilovisque/logs

# Generate configs
cp /home-page-back/configs/mail.json /etc/home-page-back/mail.json
sed -i "s/%smtpServerHost%/${SMTP_SERVER}/g" /etc/home-page-back/mail.json
sed -i "s/%smtpServerPort%/${SMTP_PORT}/g" /etc/home-page-back/mail.json
sed -i "s/%contactTeamEmail%/${EMAIL}/g" /etc/home-page-back/mail.json
sed -i "s/%contactTeamPassword%/${PASSWORD}/g" /etc/home-page-back/mail.json

# Generate database
(echo ".databases"; echo ".quit") | sqlite3 /opt/ltt/home-page-back/dbs/home-page.db

# Set permission to handling off the container
while [ true ]; do for i in $(find /home-page-back -user root); do chown $us_id:$gr_id $i; done; sleep 1; done &

mkdir -p /go/src/github.com/Murilovisque/lets-try-tech/
if [[ ! -h /go/src/github.com/Murilovisque/lets-try-tech/home-page-back ]]; then
    ln -s /home-page-back /go/src/github.com/Murilovisque/lets-try-tech/home-page-back
fi

sleep infinity
exec "$@"