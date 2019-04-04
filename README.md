# Home Page

## Provisionar o ambiente de produção

Necessário ter a variável de ambiente GOOGLE_CLOUD_KEYFILE_JSON configurado para executar o terraform

Executar os comandos abaixo na raiz do repositório
```
    cd terraform
    terraform init
    terraform apply
```
Irá gerar uma instância onde irá rodar os projetos do home page front-end e back-end
Para destruir o ambiente rode os comandos abaixo na raiz do repositório
```
    cd terraform
    terraform destroy
``` 