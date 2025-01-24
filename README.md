# ClientServerApi

Projeto de conclusão de pós-graduação (Desafio 1) 
Este projeto implementa uma API (servidor) que realiza a consulta da cotação do dolar em um serviço externo.

## Índice

1. [Descrição](#descrição)
2. [Run-Server](#run-server)
3. [Run-Client](#run-client)

## Descrição
Este projeto contem duas aplicações que realizam a consulta da cotação do dolar em um serviço externo, e está dividida em Server e Client.

### Server
A aplicação server é uma API que realiza a consulta da cotação do dolar em um serviço externo e disponibiliza a informação para o client.
Ela salva as informações em um banco de dados SQLite e disponibiliza um endpoint para consulta.

### Client
A aplicação client é responsável por realizar a consulta consumindo o endpoint https://localhost:8080/cotacao, e após receber os valores, gravar em um arquivo texto a cotação do dolar recebida.

## Run-Server
Para iniciar a aplicação server, execute o comando abaixo:
```bash
    go run server/cmd/main.go server/cmd/wire_gen.go
```
O server foi desenvolvido usando Wire para injeção de dependências, e o comando acima irá gerar o arquivo wire_gen.go que contém as dependências necessárias para a aplicação.

Tambem foi utilizado o Sqlc para gerar arquivos relacionados ao banco de dados, para gerar os arquivos execute o comando abaixo:
```bash
    sqlc generate
```
Alem destes pacotes, foram utilizados o pacote chi e viper, para expor o endpoint e controlar as variáveis de ambiente, respectivamente

### Observação: 
O arquivo de banco de dados, o exchange_db utilizado pelo server para gravar os dados acompanham o projeto.

## Run-Client

Para iniciar a aplicação client, execute o comando abaixo:
```bash
    go run client/cmd/main.go client/cmd/wire_gen.go
```
O client foi desenvolvido usando Wire para injeção de dependências, e o comando acima irá gerar o arquivo wire_gen.go que contém as dependências necessárias para a aplicação.

Foi utilizado tambem o pacote viper para controlar as variáveis de ambiente.

### Observação: 
O arquivo exchange_rate.txt utilizado pelo client para gravar a cotação do dolar acompanha o projeto com teste realizado.