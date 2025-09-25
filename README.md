# Desafio Golang: API REST de Produtos

## Visão Geral

Este projeto é uma API RESTful completa para gerenciamento de produtos, desenvolvida em Go. A aplicação segue as melhores práticas de desenvolvimento, incluindo uma arquitetura limpa em camadas, injeção de dependência, testes unitários e de integração, e documentação de API automatizada com Swagger.

## Arquitetura

O projeto utiliza uma **arquitetura em camadas** para garantir baixo acoplamento e alta coesão, tornando o código mais manutenível, escalável e testável.

-   **Handlers (Controllers)**: Responsáveis por receber as requisições HTTP, validar os dados de entrada (DTOs) e chamar a camada de serviço. Não contêm lógica de negócio.
-   **Services**: Onde reside a lógica de negócio principal da aplicação. Orquestram as operações e manipulam os dados, utilizando os repositórios.
-   **Repositories**: A única camada que interage diretamente com o banco de dados. Abstrai a lógica de acesso a dados, permitindo que o ORM (GORM) possa ser trocado facilmente no futuro.

## Funcionalidades Implementadas

A API é dividida em dois módulos principais: `auth` e `products`.

### Módulo de Autenticação (`/api/v1/auth`)

-   **`POST /auth/sign-up`**: Registra um novo usuário.
-   **`POST /auth/sign-in`**: Autentica um usuário e retorna um token JWT.
-   **`GET /auth/me`**: Retorna os dados do usuário autenticado (rota protegida).

### Módulo de Produtos (`/api/v1/products`)

-   **`POST /products`**: Cria um novo produto (rota protegida).
-   **`GET /products`**: Lista todos os produtos.
-   **`GET /products?user_id=:id`**: Filtra produtos por ID de usuário.
-   **`PUT /products/:id`**: Atualiza um produto existente (rota protegida).
-   **`DELETE /products/:id`**: Deleta um produto (rota protegida).

## Tecnologias e Ferramentas

-   **Go**: Linguagem de programação principal.
-   **Fiber**: Framework web de alta performance.
-   **GORM**: ORM para interação com o banco de dados.
-   **SQLite**: Suporte para bancos de dados SQL.
-   **Swagger (Swaggo)**: Para geração automática de documentação da API.
-   **Testify**: Suite de testes para asserções e mocks.

## Como Começar

### Pré-requisitos

-   Go (versão 1.22 ou superior)
-   Git

### 1. Clone o Repositório

```bash
git clone <URL_DO_SEU_REPOSITORIO>
cd golang-challange
```

### 2. Instale as Dependências

```bash
go mod tidy
```

### 3. Documentação da API (Swagger)

A API possui uma documentação interativa completa gerada a partir do código.
Gere os arquivos da documentação (necessário apenas após alterar os comentários godoc):

```bash
swag init -g apps/api/main.go
```

Com a aplicação rodando, acesse o seguinte URL no seu navegador:
http://localhost:3000/swagger/index.html

### 4. Como Rodar os Testes

O projeto conta com testes unitários e testes de integração.
Rodar TODOS os testes (unitários e integração):

```bash
go test ./... -v
```

### Rodar apenas os testes unitários (mais rápidos):

```bash
go test ./apps/api/modules/products/services/... -v
```

### 5. Como Rodar a Aplicação

Para rodar a aplicação, execute o seguinte comando:
```bash
go run apps/api/main.go
```
A API estará disponível em http://localhost:3000.

### 6. Collection Postman

Uma coleção do Postman está disponível no arquivo **Desafio Tech Go.postman_collection.json** para facilitar os testes manuais das rotas da API.
