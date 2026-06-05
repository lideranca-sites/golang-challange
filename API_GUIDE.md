# Guia da API - Golang Challenge

## 📋 Pré-requisitos

1. **Go instalado** (versão 1.22.6 ou superior)
2. **PostgreSQL rodando** (via Docker ou instalação local)
3. **Arquivo `.env`** configurado na raiz do projeto

## ⚙️ Configuração

### 1. Criar arquivo `.env`

Na raiz do projeto, crie um arquivo `.env` com o seguinte conteúdo:

```env
# Database Configuration
DB_HOST=localhost
DB_USER=postgres
DB_PASS=postgres
DB_NAME=postgres
DB_PORT=5432

# JWT Secret (use uma chave secreta forte em produção)
JWT_SECRET=sua-chave-secreta-jwt-aqui-mude-em-producao
```

**Importante:** Ajuste as credenciais do PostgreSQL conforme sua configuração!

### 2. Executar migrações

As migrações são executadas automaticamente quando a aplicação inicia. O banco de dados será criado/atualizado automaticamente.

## 🚀 Como Rodar a Aplicação

1. Abra um terminal na raiz do projeto
2. Execute:

```bash
cd apps/api
go run main.go
```

Ou, se estiver na raiz:

```bash
go run apps/api/main.go
```

A aplicação estará rodando em: **http://localhost:3000**

## 📡 Endpoints da API

Base URL: `{{base_url}}/api/v1`

**💡 Nota:** Para testes locais, use `http://localhost:3000` como `base_url`. Configure a variável `base_url` no Postman ou substitua `{{base_url}}` pelos comandos cURL.

---

## 🔐 Autenticação

### 1. Criar Usuário (Sign Up)

```bash
curl --location '{{base_url}}/api/v1/auth/sign-up' \
--header 'Content-Type: application/json' \
--data '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "123456"
}'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000`

**💡 Dica:** Este formato de cURL pode ser importado diretamente no Postman (Import > Raw text)

**Resposta de sucesso (201):**

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Guarde o `access_token` para usar nas requisições protegidas!**

---

### 2. Fazer Login (Sign In)

```bash
curl --location '{{base_url}}/api/v1/auth/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "email": "john@example.com",
    "password": "123456"
}'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000`

**Resposta de sucesso (200):**

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### 3. Obter Dados do Usuário Logado (Me)

```bash
curl --location '{{base_url}}/api/v1/auth/me' \
--header 'Authorization: Bearer SEU_TOKEN_AQUI'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000` e `SEU_TOKEN_AQUI` pelo `access_token` recebido do sign-up ou sign-in

**Resposta de sucesso (200):**

```json
{
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-01-06T10:00:00Z",
    "updated_at": "2026-01-06T10:00:00Z"
  }
}
```

---

## 📦 Produtos

### 4. Listar Todos os Produtos

```bash
curl --location '{{base_url}}/api/v1/products'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000`

**Resposta de sucesso (200):**

```json
{
  "products": [
    {
      "id": 1,
      "name": "Product 1",
      "price": 1000.0,
      "quantity": 10,
      "user_id": 1,
      "created_at": "2026-01-06T10:00:00Z",
      "updated_at": "2026-01-06T10:00:00Z"
    }
  ]
}
```

---

### 5. Listar Produtos por User ID

```bash
curl --location '{{base_url}}/api/v1/products?user_id=1'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000`

**Resposta de sucesso (200):**

```json
{
  "products": [
    {
      "id": 1,
      "name": "Product 1",
      "price": 1000.0,
      "quantity": 10,
      "user_id": 1,
      "created_at": "2026-01-06T10:00:00Z",
      "updated_at": "2026-01-06T10:00:00Z"
    }
  ]
}
```

---

### 6. Criar Produto ⚠️ (Requer Autenticação)

```bash
curl --location '{{base_url}}/api/v1/products' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer SEU_TOKEN_AQUI' \
--data '{
    "name": "Product 2",
    "price": 1500.50,
    "quantity": 5
}'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000` e `SEU_TOKEN_AQUI` pelo `access_token` recebido do sign-up ou sign-in

**Resposta de sucesso (201):**

```json
{
  "message": "Product created successfully",
  "product": {
    "id": 2,
    "name": "Product 2",
    "price": 1500.50,
    "quantity": 5,
    "user_id": 1,
    "created_at": "2026-01-06T10:00:00Z",
    "updated_at": "2026-01-06T10:00:00Z"
  }
}
```

**Observação:** O `user_id` é preenchido automaticamente com o ID do usuário logado (extraído do token JWT).

---

### 7. Atualizar Produto ⚠️ (Requer Autenticação)

```bash
curl --location --request PUT '{{base_url}}/api/v1/products/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer SEU_TOKEN_AQUI' \
--data '{
    "name": "Product Updated",
    "price": 2000.0,
    "quantity": 20
}'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000`, `SEU_TOKEN_AQUI` pelo `access_token` e `1` pelo ID do produto que deseja atualizar

**Resposta de sucesso (200):**

```json
{
  "message": "Product updated successfully",
  "product": {
    "id": 1,
    "name": "Product Updated",
    "price": 2000.0,
    "quantity": 20,
    "user_id": 1,
    "created_at": "2026-01-06T10:00:00Z",
    "updated_at": "2026-01-06T10:05:00Z"
  }
}
```

**Observação:** Todos os campos são opcionais no body. Você pode atualizar apenas os campos que desejar.

---

### 8. Deletar Produto ⚠️ (Requer Autenticação)

```bash
curl --location --request DELETE '{{base_url}}/api/v1/products/1' \
--header 'Authorization: Bearer SEU_TOKEN_AQUI'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000`, `SEU_TOKEN_AQUI` pelo `access_token` e `1` pelo ID do produto que deseja deletar

**Resposta de sucesso (200):**
```json
{
  "message": "Product deleted successfully"
}
```

---

## 🧪 Testando no Postman

### Opção 1: Importar Collection Completa (Recomendado)

1. **Importar Collection:**

   - Abra o Postman
   - Clique em "Import" (canto superior esquerdo)
   - Arraste o arquivo `golang-challange.postman_collection.json` ou clique em "Upload Files" e selecione o arquivo
   - A collection será importada com todas as requisições já configuradas
2. **Configurar Token:**

   - Após fazer sign-up ou sign-in, copie o `access_token` da resposta
   - Na collection, clique nos três pontos (...) ao lado do nome da collection
   - Selecione "Edit"
   - Vá na aba "Variables"
   - Cole o token na variável `token`
   - Salve
3. **Pronto!** Todas as requisições que precisam de autenticação já estão configuradas para usar `{{token}}`

### Opção 2: Importar Comandos cURL Individuais

1. No Postman, clique em "Import"
2. Selecione a aba "Raw text"
3. Cole um dos comandos cURL abaixo
4. Clique em "Import"

**Exemplo de comando cURL para importar:**

```bash
curl --location '{{base_url}}/api/v1/auth/sign-up' \
--header 'Content-Type: application/json' \
--data '{
    "name": "John Doe",
    "email": "johndoe@lidercap.com.br",
    "password": "100100"
}'
```

**💡 Para testes locais:** Substitua `{{base_url}}` por `http://localhost:3000` antes de importar

### Opção 3: Configuração Manual

#### Configuração Básica

1. **Base URL:** `http://localhost:3000/api/v1`

#### Fluxo Completo de Teste

##### Passo 1: Criar Usuário

- **Método:** POST
- **URL:** `{{base_url}}/api/v1/auth/sign-up` (ou `http://localhost:3000/api/v1/auth/sign-up` para testes locais)
- **Headers:** `Content-Type: application/json`
- **Body (raw JSON):**

```json
{
  "name": "Test User",
  "email": "test@example.com",
  "password": "123456"
}
```

- **Guarde o `access_token` da resposta!**

##### Passo 2: Configurar Token no Postman

- Crie uma variável de ambiente no Postman chamada `token`
- Cole o `access_token` recebido no Passo 1

##### Passo 3: Testar Endpoints Protegidos

- Use `Authorization: Bearer {{token}}` nos headers das requisições protegidas

---

## 📝 Variáveis de Ambiente para Postman

Crie um ambiente no Postman com as seguintes variáveis:

```
base_url = http://localhost:3000
token = (será preenchido após login)
```

**💡 Nota:** Para testes locais, use `http://localhost:3000` como `base_url`. Em produção ou outros ambientes, ajuste conforme necessário.

Depois, use nos headers e URLs:

```
Authorization: Bearer {{token}}
URL: {{base_url}}/api/v1/...
```

---

## ❌ Possíveis Erros

### 401 Unauthorized

- Token JWT inválido ou expirado
- Token não enviado no header `Authorization: Bearer`
- Credenciais inválidas no sign-in (usuário não encontrado ou senha incorreta)

**Exemplo de resposta de credenciais inválidas:**

```json
{
  "error": "Invalid credentials"
}
```

### 400 Bad Request

- Campos obrigatórios faltando
- Formato de email inválido
- Dados no formato incorreto

**Exemplo de resposta de erro de validação:**

```json
{
  "field": "Email",
  "tag": "email",
  "message": "The Email field must be a valid email"
}
```

**Exemplo de resposta de campo obrigatório:**

```json
{
  "field": "Password",
  "tag": "required",
  "message": "The Password field is required"
}
```

### 404 Not Found

- Produto com ID especificado não existe
- Endpoint não encontrado

### 500 Internal Server Error

- Erro de conexão com banco de dados
- Erro interno do servidor

---

## 🔍 Validações

### Sign Up / Sign In

- `name`: obrigatório (apenas Sign Up)
- `email`: obrigatório, formato válido de email
- `password`: obrigatório

### Criar Produto

- `name`: obrigatório (string)
- `price`: obrigatório (float64)
- `quantity`: obrigatório (int)

### Atualizar Produto

- Todos os campos são opcionais
- `name`: string (opcional)
- `price`: float64 (opcional)
- `quantity`: int (opcional)

---

## ✅ Checklist de Testes

- [ ] Criar usuário (sign-up)
- [ ] Fazer login (sign-in)
- [ ] Obter dados do usuário (me) - com token
- [ ] Listar todos os produtos (sem token)
- [ ] Listar produtos por user_id (sem token)
- [ ] Criar produto (com token)
- [ ] Atualizar produto (com token)
- [ ] Deletar produto (com token)
- [ ] Testar acesso sem token (deve retornar 401)

---

## 🎯 Exemplos Práticos

### Exemplo 1: Criar usuário e produto

```bash
# 1. Criar usuário
curl --location 'http://localhost:3000/api/v1/auth/sign-up' \
--header 'Content-Type: application/json' \
--data '{"name": "Maria Silva", "email": "maria@example.com", "password": "senha123"}'

# Resposta: {"access_token": "eyJhbGc..."}

# 2. Criar produto (usar o token recebido)
curl --location 'http://localhost:3000/api/v1/products' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGc...' \
--data '{"name": "Notebook", "price": 3500.00, "quantity": 3}'
```

### Exemplo 2: Listar produtos do usuário

```bash
# 1. Fazer login
curl --location 'http://localhost:3000/api/v1/auth/sign-in' \
--header 'Content-Type: application/json' \
--data '{"email": "maria@example.com", "password": "senha123"}'

# 2. Listar produtos do usuário (substitua 1 pelo ID do usuário)
curl --location 'http://localhost:3000/api/v1/products?user_id=1'
```

---

## 📦 Arquivos de Importação para Postman

### Collection Completa

- **Arquivo:** `golang-challange.postman_collection.json`
- **Como usar:** Import > Upload Files > Selecione o arquivo
- **Vantagem:** Todas as requisições já configuradas, incluindo variáveis de token

### Comandos cURL Individuais

- **Arquivo:** `postman-curl-commands.txt`
- **Como usar:** Import > Raw text > Cole o comando desejado
- **Vantagem:** Importar apenas as requisições que você precisa
