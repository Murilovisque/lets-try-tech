#!/bin/bash

npm install
# Set permission to handling off the container
while [ true ]; do for i in $(find /home-page-front -user root); do chown $us_id:$gr_id $i; done; sleep 1; done &

/bin/bash