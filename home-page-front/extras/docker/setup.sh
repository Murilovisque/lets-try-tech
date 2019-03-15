#!/bin/bash

npm install
# Set permission to handling off the container
while [ true ]; do chown -R $us_id:$gr_id /home-page-front; sleep 1; done &

/bin/bash