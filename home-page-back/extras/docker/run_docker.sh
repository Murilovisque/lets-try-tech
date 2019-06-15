#!/bin/bash

for i in SMTP_SERVER SMTP_PORT EMAIL PASSWORD; do
    if [[ -z $(eval echo '$'$i) ]]; then
        read -s -p "$i: " contentVar
        export `echo $i`=`echo $contentVar`
        echo
    fi
done

docker stop --time=1 home-page-back
docker build -t home-page-back extras/docker/
docker run -d -v $(pwd):/home-page-back -v go-source:/go/src -e us_id=`id -u` -e gr_id=`id -g` -e SMTP_SERVER=${SMTP_SERVER} -e SMTP_PORT=${SMTP_PORT} -e EMAIL=${EMAIL} -e PASSWORD=${PASSWORD} --rm --net=host --name home-page-back home-page-back