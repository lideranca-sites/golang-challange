# Desafio Golang: Cria um CRUD de Produtos

## Objetivo Cumprido

Este projeto implementa uma criação de API REST para o gerenciamento de produtos. O foco foi a criação de um módulo `products` com funcionalidades completas de CRUD (Create, Read, Update, Delete).

## Funcionalidades Implementadas

O módulo de produtos `/api/v1/products` inclui os seguintes endpoints:

* **`POST /products`**: Cria um novo produto. (Rota protegida por JWT)
* **`GET /products`**: Lista todos os produtos.
* **`GET /products?user_id=:id`**: Filtra os produtos por ID de usuário.
* **`PUT /products/:id`**: Atualiza um produto existente. (Rota protegida por JWT)
* **`DELETE /products/:id`**: Deleta um produto. (Rota protegida por JWT)

A implementação foi desenvolvida seguindo a abordagem de TDD (Test-Driven Development), garantindo que todas as funcionalidades fossem cobertas pelos testes de integração contidos em `product_test.go`.

## Tecnologias

* **Go**
* **Fiber** (Framework Web)
* **GORM** (ORM para banco de dados)
* **SQLite** (SQLite)

## Como Rodar os Testes

Para verificar a funcionalidade do módulo de produtos, execute:

```bash
go test product_test.go -v
```

## Como Rodar a aplicação

Para rodar a aplicação, execute:

```bash
go run apps/api/main.go
```

## Collection do Postman para teste das rotas e banco de dados

Json da collection para testes:

Arquivo: **Desafio Tech Go.postman_collection.json**