Fruit API - Desafio Go Júnior

API desenvolvida em Go como parte de um desafio técnico para vaga Júnior Backend. A aplicação consulta frutas de uma API pública, armazena localmente e gera relatórios baseados na quantidade de açúcar.


---

Tecnologias utilizadas

Go

Gin (framework HTTP)

GORM (ORM)

SQLite (banco de dados local)

robfig/cron (agendador de tarefas)



---

Funcionalidades

Parte 1: Crawler de Frutas

Roda automaticamente todos os dias à meia-noite (via cron)

Consulta a API pública Fruityvice

Armazena e atualiza as frutas no banco local


Parte 2: Relatório de Açúcar

Endpoint para gerar um relatório com frutas agrupadas por quantidade de açúcar

Frutas com açúcar >= 10 são consideradas high_sugar

Frutas com açúcar < 10 são consideradas low_sugar



---

Como rodar o projeto localmente

Pré-requisitos:

Go instalado (download aqui)


Passos:

1. Clone o repositório:



git clone https://github.com/seu-usuario/seu-repo.git
cd seu-repo

2. Baixe as dependências:



go mod tidy

3. Rode a aplicação:



go run ./go-fruit-api/cmd/api

> O servidor estará disponível em: http://localhost:8080




---

Rotas disponíveis

GET /api/fruits/load

Carrega as frutas da API pública e armazena/atualiza no banco local.

Resposta esperada:

{
  "message": "Frutas carregadas com sucesso"
}

GET /api/fruits/report-sugar

Retorna um relatório agrupando frutas por nível de açúcar.

Resposta exemplo:

{
  "high_sugar": [
    {
      "id": 67,
      "name": "Lychee",
    }
  ],
  "low_sugar": [
    {
      "id": 6,
      "name": "Banana",
    }
  ],
  "total_high_sugar": 1,
  "total_low_sugar": 1
}


---

Estrutura do projeto

go-fruit-api/
├── cmd/api/            # Arquivo main.go e inicialização do servidor
├── controllers/        # Handlers das rotas
├── crawler/            # Função que consome a API externa
├── database/           # Inicialização do banco SQLite
├── models/             # Structs GORM
├── routes/             # Registro das rotas
├── scheduler/          # Tarefa agendada (cron)
├── services/           # Lógica de negócios
├── go.mod / go.sum     # Gerenciador de dependências


---

Observações finais

O banco de dados SQLite (fruits.db) não é versionado.

A tarefa agendada via cron executa diariamente, mas pode ser ajustada para testes com @every 1m.