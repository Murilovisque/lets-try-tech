#!/bin/bash

npm install
# Set permission to handling off the container
while [ true ]; do for i in $(find /home-page-front -user root); do chown $us_id:$gr_id $i; done; sleep 1; done &

rm -rf /etc/nginx/nginx.conf
ln -s /home-page-front/extras/nginx/nginx.conf /etc/nginx/nginx.conf
ln -s /home-page-front/extras/nginx/home-page /etc/nginx/sites-available/home-page
rm -rf /etc/nginx/sites-enabled/*
ln -s /etc/nginx/sites-available/home-page /etc/nginx/sites-enabled/home-page
service nginx start

sleep infinity