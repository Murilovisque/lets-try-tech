#!/bin/bash

cp /home-page-back/configs/mail.json /opt/ltt/home-page-back/configs/mail.json

sed -i "s/%smtpServerHost%/${SMTP_SERVER}/g" /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%smtpServerPort%/${SMTP_PORT}/g" /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%contactTeamEmail%/${EMAIL}/g" /opt/ltt/home-page-back/configs/mail.json
sed -i "s/%contactTeamPassword%/${PASSWORD}/g" /opt/ltt/home-page-back/configs/mail.json

export PASSWORD=""

# Set permission to handling off the container
while [ true ]; do for i in $(find /home-page-back -user root); do chown $us_id:$gr_id $i; done; sleep 1; done &

nc -l 666
exec "$@"