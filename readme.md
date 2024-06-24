# Desafio de Aplicação Go: Servidor e Cliente

Este repositório contém uma aplicação Go composta por dois sistemas: um servidor (`server/app/main.go`) e um cliente (`client/app/main.go`). A aplicação é projetada para demonstrar o uso de webserver HTTP, contextos, banco de dados e manipulação de arquivos com Go.

## Descrição do Desafio

Neste desafio, aplicamos os conceitos aprendidos para:

- Consumir uma API externa para obter a cotação do dólar.
- Registrar cotações em um banco de dados SQLite.
- Implementar timeouts com o package `context` em diferentes partes do sistema.
- Salvar cotações em um arquivo de texto(txt).

### Requisitos

1. O `client/app/main.go` deve fazer uma requisição HTTP para o `server/app/main.go`, solicitando a cotação do dólar.
2. O `server/app/main.go` deve consumir a API de câmbio de Dólar e Real no endereço: [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL).
3. O `server/app/main.go` deve retornar o resultado para o cliente no formato JSON.
4. O `server/app/main.go` deve registrar a cotação recebida no banco de dados SQLite, com um timeout máximo de 10ms para persistir os dados.
5. O `client/app/main.go` deve receber apenas o valor atual do câmbio (campo "bid" do JSON).
6. O `client/app/main.go` deve salvar a cotação atual em um arquivo `cotacao.txt` no formato: `Dólar: {valor}`.
7. O `client/app/main.go` deve ter um timeout máximo de 300ms para receber o resultado do `server.go`.
8. Todos os contextos devem registrar erros nos logs caso o tempo de execução seja insuficiente.

### Endpoints e Porta

- Endpoint necessário: `/cotacao`
- Endpoint para consultar dados salvos no SQLite: `/getallcotacao`
- Porta do servidor HTTP: `8080`

## Instruções para Configuração e Execução

### Pré-requisitos

- Go instalado na máquina.

### Passo a Passo

-> cd server/app
-> go mod tidy
-> go run .

-> cd client/app
-> go run .

.
├── client
│   └── app
│       └── client.go
├── server
│   └── app
│       └── server.go
├── README.md
└── go.mod

Feito com ❤️ por Vinicius