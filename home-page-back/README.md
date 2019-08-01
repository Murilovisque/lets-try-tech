# Home Page Back

## Ferramentas necessárias
* Docker -> https://docs.docker.com/install/

## Criar o ambiente via docker

Executar o comando abaixo para iniciar o container em background na raiz do projeto home-page-back e acessá-lo
```
./extras/docker/run_docker.sh
```

## Compilar e gerar o executável do projeto

Executar o comando abaixo no container docker para iniciar o projeto
```
go run /home-page-back/cmd/home-page-back/main.go
```

Executar os comandos abaixo no container docker. Irá gerar um arquivo executável de nome 'home-page'
```
cd /home-page-back/cmd/home-page
go build
```

## Executar os testes
```
cd /home-page-back
go test ./...
```

## Estrutura do projeto

A estrutura do projeto se baseia nos padrões abaixo

* https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
* https://github.com/golang-standards/project-layout

## Gerar um pacote debian do projeto. Deve ser executado dentro do container do Docker. O arquivo será gerado na pasta 'build/package/debian/target'
```
cd /home-page-back/build/package/debian/
./build.sh
```