# Home Page Front

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 7.3.4.

## Criar o ambiente via docker

Executar os comandos abaixo para iniciar o container em background
```
docker build -t home-page-front extras/docker/
docker run -d -v $(pwd):/home-page-front -e us_id=`id -u` -e gr_id=`id -g` --rm --net=host --name home-page-front home-page-front
```

## Acessar o container docker
```
docker exec -it home-page-front bash
```

## Executar o servidor em modo desenvolvimento

Dentro do container execute o comando abaixo para disponibilizar a aplicação no link `http://localhost:4200/`. Alterações irá automaticamente regarregar o projeto
```
ng serve
```

## Gerar o pacote do projeto

Executar o comando abaixo no container docker. Irá gerar os arquivos na pasta 'dist/'
```
ng build --prod
```

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI README](https://github.com/angular/angular-cli/blob/master/README.md).
