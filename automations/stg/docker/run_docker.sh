#!/bin/bash

for i in SMTP_SERVER SMTP_PORT EMAIL PASSWORD; do
    if [[ -z $(eval echo '$'$i) ]]; then
        read -s -p "$i: " contentVar
        export `echo $i`=`echo $contentVar`
        echo
    fi
done

docker stop --time=1 home-page-stg
docker build -t home-page-stg automations/stg/docker/
docker run -d --net=host --rm -v $(pwd)/automations/stg/docker:/home-page-stg/docker -v $(pwd)/automations/stg/nginx:/home-page-stg/nginx -v $(pwd)/home-page-front/dist/home-page-front:/opt/ltt/home-page-front -v $(pwd)/home-page-back/build/package/debian/target:/home-page-stg/home-page-back/debian -v $(pwd)/home-page-back/configs:/home-page-stg/home-page-back/configs --name home-page-stg home-page-stg