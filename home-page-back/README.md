# Home Page Back

## Ferramentas necessárias
* Docker -> https://docs.docker.com/install/

## Criar o ambiente via docker

Executar o comando abaixo para iniciar o container em background na raiz do projeto home-page-back
```
./extras/docker/run_docker.sh
```

## Acessar o container docker
```
docker exec -it home-page-back bash
```

## Estrutra do projeto

A estrutura do projeto se baseia nos padrões abaixo

* https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
* https://github.com/golang-standards/project-layout