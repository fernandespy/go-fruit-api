# Desafio Go

## Especificação:

Criar uma API em GO (Usar framework Gin https://github.com/gin-gonic/gin) com GORM (https://gorm.io/), pode ser banco SQLite para simplificar

## Desafio

**Parte 1** -  Crawller que roda todo dia a meia noite e carrega os dados das frutas, lembrando que as frutas podem sofrer alteração de valores, então atualizar os dados, salvar os dados no banco de dados

**Parte 2** -  Criar endpoint `/api/fruits/report-sugar` que será responsável por extrair o relatório das frutas sobre sua nutrição:

Fonte de dados:
- https://www.fruityvice.com/#3
- https://www.fruityvice.com/doc/index.html
- https://www.fruityvice.com/api/fruit/all

Regra: 
- High Sugar: Maior ou igual 10
- Low Sugar: Menor que 10

Exemplo de retorno

```
{
	"high_sugar": [
		{
			 "name": "Lychee",
			 "id": 67
		}
	],
	"low_sugar": [
		{
			 "name": "Lychee",
			 "id": 67
		}
	],
	"total_high_sugar": 1,
	"total_low_sugar": 1,
}
```




