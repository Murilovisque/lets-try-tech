# Home Page Back

## Ferramentas necessárias
* Docker -> https://docs.docker.com/install/

Este projeto foi gerado com o angular-cli [Angular CLI](https://github.com/angular/angular-cli) na versão 7.3.4.

## Criar o ambiente via docker

Executar os comandos abaixo para iniciar o container em background
```
read -s -p "Pass: " PASSWORD
docker build -t home-page-back extras/docker/
docker run -d -v $(pwd):/home-page-back -e us_id=`id -u` -e gr_id=`id -g` -e SMTP_SERVER=${SMTP_SERVER} -e SMTP_PORT=${SMTP_PORT} -e EMAIL=${EMAIL} -e PASSWORD=${PASSWORD} --rm --net=host --name home-page-back home-page-back
```

## Acessar o container docker
```
docker exec -it home-page-back bash
```