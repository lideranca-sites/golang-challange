# Desafio Golang

## Objetivo

O objetivo deste desafio é avaliar a curva de aprendizado do nosso time em migrar de uma typescript de programação para Golang.

## Desafio

O desafio consiste em criar uma API REST que seja capaz de realizar as operações de CRUD (Create, Read, Update, Delete) de um recurso chamado `Product`.

### Requisitos

- O recurso `Product` deve possuir os seguintes campos:
  - `id` (int)
  - `name` (string)
  - `price` (float64)
  - `quantity` (int)
  - `user_id` (int)

**A API deve possuir os seguintes endpoints:**

- `GET /products`: Retorna a lista de todos os produtos cadastrados.
- `GET /products/:id`: Retorna os detalhes de um produto específico.
- `GET /users/:id/products`: Retorna a lista de produtos de um usuário específico.
- `POST /users/:id/products`: Cria um novo produto para um usuário específico. (Apenas o proprio usuário autenticado pode criar um produto para si mesmo)
- `PUT /products/:id`: Atualiza os dados de um produto específico. (Apenas o proprio usuário autenticado pode atualizar um produto para si mesmo)
- `DELETE /products/:id`: Deleta um produto específico. (Apenas o proprio usuário autenticado pode deletar um produto para si mesmo)

### Observações

Todas as rotas devem ser autenticadas, exceto a rota de listagem de produtos.

