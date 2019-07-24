# Home Page Front

## Ferramentas necessárias
* Docker -> https://docs.docker.com/install/

Este projeto foi gerado com o angular-cli [Angular CLI](https://github.com/angular/angular-cli) na versão 7.3.4.

## Criar o ambiente via docker

Executar os comandos abaixo para iniciar o container em background
```
./extras/docker/run_docker.sh
```

## Acessar o container docker
```
docker exec -it home-page-front bash
```

## Executar o servidor em modo desenvolvimento

Dentro do container execute o comando abaixo para disponibilizar a aplicação no link `http://localhost:4200/`. Alterações irão automaticamente recarregar o projeto
```
ng serve
```

## Dicas de desenvolvimento

Comando 'generate ...' irá gerar o componente desejado 
```
ng generate component component-name
ng generate directive|pipe|service|class|guard|interface|enum|module
```

### Executar testes unitários

Testes unitários via [Karma](https://karma-runner.github.io).
```
ng test
```

### Rodando testes end-to-end

Testes end-to-end via tests via [Protractor](http://www.protractortest.org/).
```
ng e2e
```

## Gerar o pacote do projeto

Executar o comando abaixo no container docker. Irá gerar os arquivos na pasta 'dist/'
```
ng build --prod
```