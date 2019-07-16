FROM debian:stretch
RUN apt update
RUN apt install nginx -y
RUN apt install vim -y
RUN apt install sqlite3 -y
RUN apt install procps -y
RUN apt install logrotate -y

RUN mkdir -p /home-page-stg /opt/ltt/
RUN mkdir -p /var/log/home-page-back/archive
RUN mkdir -p /var/lib/ltt/home-page-back/dbs
RUN mkdir -p /home-page-stg/home-page-back/

EXPOSE 80
ENTRYPOINT /bin/bash -c "chmod +x /home-page-stg/docker/entrypoint.sh && /home-page-stg/docker/entrypoint.sh"