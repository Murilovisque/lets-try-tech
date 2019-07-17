#!/bin/bash

rm -rf /etc/nginx/nginx.conf
ln -s /home-page-stg/nginx/nginx.conf /etc/nginx/nginx.conf
ln -s /home-page-stg/nginx/sites-available/home-page /etc/nginx/sites-available/home-page
rm -rf /etc/nginx/sites-enabled/*
ln -s /etc/nginx/sites-available/home-page /etc/nginx/sites-enabled/home-page

service nginx start

dpkg -i /home-page-stg/home-page-back/debian/home-page-back.deb

# Copy database
cp /home-page-stg/home-page-back/debian/home-page.db /var/lib/ltt/home-page-back/dbs/

# Copy configs
cp /home-page-stg/home-page-back/debian/mail.json /etc/home-page-back/mail.json

service home-page-back start

sleep infinity
exec "$@"