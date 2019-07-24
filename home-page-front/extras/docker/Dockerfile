FROM node:10.16.0-stretch
RUN apt update
RUN apt install nginx -y
EXPOSE 4200
EXPOSE 80
RUN echo 'export PATH=$PATH:/home-page-front/node_modules/.bin/' >> /root/.bashrc

WORKDIR /home-page-front
ENTRYPOINT /bin/bash -c "chmod +x /home-page-front/extras/docker/entrypoint.sh && /home-page-front/extras/docker/entrypoint.sh"
