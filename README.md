# Home Page

## Ferramentas necessárias
* Docker -> https://docs.docker.com/install/
* Ansible -> https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html
* Terraform -> https://learn.hashicorp.com/terraform/getting-started/install.html
* Terraform-inventory -> https://github.com/adammck/terraform-inventory/releases

## Ambiente de stage via docker

Executar os passos da seção "Gerar o pacote do projeto" descritos no README.md do projeto **home-page-front** afim de disponibiliza-lo no ambiente stage

Executar os comandos abaixo na raiz do repositório para iniciar o container em background. O projeto estará disponível acessando o link http://localhost
```
./automations/stg/docker/run_docker.sh
```

## Acessar o container docker
```
docker exec -it home-page-stg bash
```

## Provisionar o ambiente de produção

### Criando o infraestrutura

Necessário ter a variável de ambiente GOOGLE_CLOUD_KEYFILE_JSON configurado para executar o terraform

Executar os comandos abaixo na raiz do repositório para gerar uma instância onde irá rodar os projetos do home page front-end e back-end
```
cd automations/prod/terraform
terraform init
terraform apply
```

### Configurando os hosts da infraestrutura

Necessário ter configurado o acesso via SSH para o host da infraestrutura criado no passo anterior

Execute os comandos abaixo na raiz do repositório para configurar as instâncias para execução do projeto. É necessário ter a variável ${EMAIL} preenchida
```
cd automations/prod/ansible
TEMP_DIR=`mktemp -d`; terraform-inventory --inventory ../terraform/ > $TEMP_DIR/inventory; ansible-playbook -i $TEMP_DIR/inventory main.yml --extra-vars "email=${EMAIL}"; rm -rf $TEMP_DIR;
```

Caso queira rodar somente uma parte da automação de configuração das instâncias basta informar quais tags devem ser executadas. É necessário ter a variável ${EMAIL} preenchida:

* nginx -> Instala e configuração do nginx
* nginx-certificate -> Instala o certbot e os certificados do nginx
* home-page-front -> Instala e configuração do projeto home-page-front

```
cd automations/prod/ansible
TEMP_DIR=`mktemp -d`; terraform-inventory --inventory ../terraform/ > $TEMP_DIR/inventory; ansible-playbook -i $TEMP_DIR/inventory --tags "nginx" main.yml --extra-vars "email=${EMAIL}"; rm -rf $TEMP_DIR;
```

### Destruindo a infraestrutura

Execute os comandos abaixo na raiz do repositório
```
cd automations/prod/terraform
terraform destroy
``` 