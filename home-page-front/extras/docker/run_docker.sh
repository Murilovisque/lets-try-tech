#!/bin/bash

docker stop --time=1 home-page-front
docker build -t home-page-front extras/docker/
docker run -d -v $(pwd):/home-page-front -e us_id=`id -u` -e gr_id=`id -g` --rm --net=host --name home-page-front home-page-front