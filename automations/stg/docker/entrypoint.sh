#!/bin/bash

rm -rf /etc/nginx/nginx.conf
ln -s /home-page-stg/nginx/nginx.conf /etc/nginx/nginx.conf
ln -s /home-page-stg/nginx/sites-available/home-page /etc/nginx/sites-available/home-page
rm -rf /etc/nginx/sites-enabled/*
ln -s /etc/nginx/sites-available/home-page /etc/nginx/sites-enabled/home-page

service nginx start

nc -l 666
exec "$@"