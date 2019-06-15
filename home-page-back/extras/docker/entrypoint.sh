#!/bin/bash

# Get all project dependencies
go get ./...

# Generate configs
cp /home-page-back/configs/mail.json /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%smtpServerHost%/${SMTP_SERVER}/g" /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%smtpServerPort%/${SMTP_PORT}/g" /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%contactTeamEmail%/${EMAIL}/g" /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%contactTeamPassword%/${PASSWORD}/g" /opt/ltt/home-page-back/configs/mail.json
export PASSWORD=""

# Generate database
(echo ".databases"; echo ".quit") | sqlite3 /opt/ltt/home-page-back/dbs/home-page.db

# Set permission to handling off the container
while [ true ]; do for i in $(find /home-page-back -user root); do chown $us_id:$gr_id $i; done; sleep 1; done &

# Build function
echo 'function build-app() {
    local currentFolder=$(pwd)    
    cd `find /home-page-back -name main.go -printf "%h\n"`
    go build
    cd $currentFolder
}' >> /root/.bashrc

mkdir -p /go/src/github.com/Murilovisque/lets-try-tech/

if [[ ! -h /go/src/github.com/Murilovisque/lets-try-tech/home-page-back ]]; then
    ln -s /home-page-back /go/src/github.com/Murilovisque/lets-try-tech/home-page-back
fi

sleep infinity
exec "$@"