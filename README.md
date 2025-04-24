# Fruit API - Desafio Go Júnior

## API desenvolvida em Go como parte de um desafio técnico para vaga Júnior Backend. A aplicação consulta frutas de uma API pública, armazena localmente e gera relatórios baseados na quantidade de açúcar.


---

### Tecnologias utilizadas
```
- Go

- Gin (framework HTTP)

- GORM (ORM)

- SQLite (banco de dados local)

- robfig/cron (agendador de tarefas)
```


---

### Funcionalidades

Parte 1: Crawler de Frutas

- Roda automaticamente todos os dias à meia-noite (via cron)

- Consulta a API pública Fruityvice

- Armazena e atualiza as frutas no banco local


Parte 2: Relatório de Açúcar

- Endpoint para gerar um relatório com frutas agrupadas por quantidade de açúcar

- Frutas com açúcar >= 10 são consideradas high_sugar

- Frutas com açúcar < 10 são consideradas low_sugar



---

### Como rodar o projeto localmente

Pré-requisitos:

Go instalado (https://go.dev/dl/)

Passos:

1. Clone o repositório:


git clone https://github.com/fernandespy/go-fruit-api.git

2. Baixe as dependências:

`go mod tidy`

3. Rode a aplicação:

`go run ./cmd/api`

Possível erro encontrado aqui é quando rodar o projeto e o receber o erro de **CGO_ENABLED=0* pois a versão do Go pode não ser suportada pelo go-sqlite3

Solução:
- primeiro set o CGO para 1 usando:
> $env:CGO_ENABLED=1
- Depois baixe o mingw para auxiliar no compilador C usando:
> 1: Set-ExecutionPolicy RemoteSigned -Scope CurrentUser irm get.scoop.sh | iex
> 2: scoop install mingw
- Ao finalizar a instalação, execute novamente go run ./cmd/api e aplicação deve rodar normalmente.

> O servidor estará disponível em: http://localhost:8080


---

### Rotas disponíveis

 *GET /api/fruits/load*

Carrega as frutas da API pública e armazena/atualiza no banco local.

Resposta esperada:

```json
{
  "message": "Fruits loaded successfuly"
}
```

*GET /api/fruits/report-sugar*

Retorna um relatório agrupando frutas por nível de açúcar.

Resposta exemplo:

```json
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
```

*GET /api/fruits/*

Retorna todas as frutas existentes no db

Resposta exemplo

```json
{
    "Fruits": [
        {
            "ID": 2, //ID único gerado pelo proprio DB
            "CreatedAt": "2025-04-24T14:52:12.0811751-03:00",
            "UpdatedAt": "2025-04-24T14:52:12.0811751-03:00",
            "DeletedAt": null,
            "id": 3, // ID respectivo da fruta disponivel na API, e ID usado nas rotas get e delete
            "name": "Strawberry",
            "family": "Rosaceae",
            "genus": "Fragaria",
            "order": "Rosales",
            "carbohydrates": 5.5,
            "protein": 0.8,
            "fat": 0.4,
            "calories": 29,
            "sugar": 5.4
        },
        {
            "ID": 4, //ID único gerado pelo proprio DB
            "CreatedAt": "2025-04-24T14:52:12.0982658-03:00",
            "UpdatedAt": "2025-04-24T14:52:12.0982658-03:00",
            "DeletedAt": null,
            "id": 5, // ID respectivo da fruta disponivel na API, e ID usado nas rotas get e delete
            "name": "Tomato",
            "family": "Solanaceae",
            "genus": "Solanum",
            "order": "Solanales",
            "carbohydrates": 3.9,
            "protein": 0.9,
            "fat": 0.2,
            "calories": 74,
            "sugar": 2.6
        },
```

*GET /api/fruits/:fruit_id*
> exemplo: [http:](http://localhost:8080/api/fruits/52)

Retorna a fruta que possui o respectivo ID, mas primeira checa se a fruta teve alguma deleção aplicada pelo campo deleted_at no DB, se não retorna a mesma

Retorno da fruta exemplo:

```json
{
    "Fruit": {
        "ID": 2,
        "CreatedAt": "2025-04-24T14:52:12.0811751-03:00",
        "UpdatedAt": "2025-04-24T14:52:12.0811751-03:00",
        "DeletedAt": null,
        "id": 3,
        "name": "Strawberry",
        "family": "Rosaceae",
        "genus": "Fragaria",
        "order": "Rosales",
        "carbohydrates": 5.5,
        "protein": 0.8,
        "fat": 0.4,
        "calories": 29,
        "sugar": 5.4
    }
}
```

Retorno caso a fruta seja deletada:

```json
{
    "Fruit was deleted_at": "2025-04-24T15:05:44.696144-03:00"
}
```
*GET /api/fruits/delete/:fruit_id*
> exemplo: [http:](http://localhost:8080/api/fruits/delete/52)

Realiza a deleção da fruta especificada direto no DB

Retorno esperado:

```json
{
    "message": "Fruit deleted successfully"
}
```

Possivel resposta:

```json
{
    "error": "Fruit not found"
}
```

Ou:

```json
{
    "Fruit already deleted:": "2025-04-24T15:34:47.8652213-03:00"
}
```

*Observação:*

Não criei um method de POST, pois como pegamos os dados da API, adicionar frutas próprias seria desnecessário. 

---

### Estrutura do projeto

```
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
```

---

### Observações finais

O banco de dados SQLite (fruits.db) não é versionado.

A tarefa agendada via cron executa diariamente, mas pode ser ajustada para testes com @every 1m.
