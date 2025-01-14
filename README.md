# Desafio Golang

//warning
> [!CAUTION]
> Não acesse a branch de outras pessoas. Crie sua própria branch e faça o PR para a branch master.
>
> Roubar e coisa de trouxa. Não seja trouxa.

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

- `GET /products?user_id=x`: Deve retornar a lista de todos os produtos cadastrados ou filtrar por `user_id` caso seja passado como query param.
- `POST /products`: Deve criar um novo produto.
- `PUT /products/:id`: Deve atualizar o produto com o `id` especificado.
- `DELETE /products/:id`: Deve deletar o produto com o `id` especificado.

### Observações

Todas as rotas exceto de consulta (`GET`) devem ser protegidas por autenticação. A autenticação deve ser feita através de um token JWT.

### Requisitos técnicos

- Criar um modulo products dentro da pasta `/apps/api/modules` para organizar o código.
- Adicionar campos ao modelo `Product`.
- Criar uma feature por arquivo.
- Criar um arquivo de rotas para o módulo `Product`.
- Adicionar as rotas do módulo `Product` ao arquivo de rotas principal.

### Como começar

- Crie uma branch a partir da branch `master` com o nome `<seu-nome>`.
- Desenvolva sua solucao com TDD (Testes ja estão criados).
- Ao finalizar, abra um PR para a branch `master` e solicite a revisão do seu código.

### Dicas

Para rodar os testes, execute o comando `go test products_test.go`. Utilize a flag `-v` para ver o output dos testes com mais detalhes.
