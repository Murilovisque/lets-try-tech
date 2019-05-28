FROM debian:stretch
RUN apt update
RUN apt install nginx -y
RUN apt install vim -y
RUN apt install netcat -y

RUN mkdir -p /home-page-stg /opt/ltt/

EXPOSE 80
ENTRYPOINT /bin/bash -c "chmod +x /home-page-stg/docker/entrypoint.sh && /home-page-stg/docker/entrypoint.sh"