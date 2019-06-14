FROM golang:stretch
RUN apt update
RUN apt install netcat -y
RUN apt install sqlite3 -y

RUN mkdir -p /opt/ltt/home-page-back/configs/
RUN mkdir -p /opt/ltt/home-page-back/dbs/
WORKDIR /home-page-back

ENTRYPOINT /bin/bash -c "chmod +x /home-page-back/extras/docker/entrypoint.sh && /home-page-back/extras/docker/entrypoint.sh"