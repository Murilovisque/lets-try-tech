#!/bin/bash

rm -rf /etc/nginx/nginx.conf
ln -s /home-page-stg/nginx/nginx.conf /etc/nginx/nginx.conf
ln -s /home-page-stg/nginx/sites-available/home-page /etc/nginx/sites-available/home-page
rm -rf /etc/nginx/sites-enabled/*
ln -s /etc/nginx/sites-available/home-page /etc/nginx/sites-enabled/home-page

service nginx start

(echo ".databases"; echo ".quit") | sqlite3 /var/lib/ltt/home-page-back/dbs/home-page.db
dpkg -i /home-page-stg/home-page-back/debian/home-page-back.deb

cp /home-page-stg/home-page-back/configs/mail.json /etc/home-page-back/mail.json
sed -i "s/%smtpServerHost%/${SMTP_SERVER}/g" /etc/home-page-back/mail.json
sed -i "s/%smtpServerPort%/${SMTP_PORT}/g" /etc/home-page-back/mail.json
sed -i "s/%contactTeamEmail%/${EMAIL}/g" /etc/home-page-back/mail.json
sed -i "s/%contactTeamPassword%/${PASSWORD}/g" /etc/home-page-back/mail.json

service home-page-back start

sleep infinity
exec "$@"