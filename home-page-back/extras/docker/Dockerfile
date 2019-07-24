FROM golang:stretch
RUN apt update
RUN apt install sqlite3 -y

RUN mkdir -p /opt/ltt/home-page-back/configs
RUN mkdir -p /var/lib/ltt/home-page-back/dbs
RUN mkdir -p /etc/home-page-back
RUN mkdir -p /var/log/home-page-back/archive
RUN rm -r /go/src/
EXPOSE 8080
WORKDIR /home-page-back

ENTRYPOINT /bin/bash -c "chmod +x /home-page-back/extras/docker/entrypoint.sh && /home-page-back/extras/docker/entrypoint.sh"