FROM node:10.15.2
RUN npm install -g @angular/cli
EXPOSE 4200
WORKDIR /home-page-front
ENTRYPOINT /bin/bash -c "chmod +x /home-page-front/extras/docker/setup.sh && /home-page-front/extras/docker/setup.sh"
