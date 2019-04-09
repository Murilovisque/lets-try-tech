# Home Page

## Ambiente de stagging via docker

Necessário executar os passos descritos no README.md do projeto home-page-front para geração do pacote

Execute os comandos abaixo para iniciar o container em background. O projeto estará disponível acessando o link http://localhost
```
docker build -t home-page-stag automations/docker/
docker run -d --net=host --rm -v $(pwd)/automations/docker:/home-page-stag/docker -v $(pwd)/automations/nginx:/home-page-stag/nginx -v $(pwd)/home-page-front/dist/home-page-front:/opt/ltt/home-page-front --name home-page-stag home-page-stag
```

## Acessar o container docker
```
docker exec -it home-page-stag bash
```

## Provisionar o ambiente de produção

Necessário ter a variável de ambiente GOOGLE_CLOUD_KEYFILE_JSON configurado para executar o terraform

Executar os comandos abaixo na raiz do repositório
```
    cd automations/terraform
    terraform init
    terraform apply
```
Irá gerar uma instância onde irá rodar os projetos do home page front-end e back-end
Para destruir o ambiente rode os comandos abaixo na raiz do repositório
```
    cd terraform
    terraform destroy
``` 