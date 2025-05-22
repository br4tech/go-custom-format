# 📦 GoCustomFormat-ProductAPI

Uma aplicação de estudo em Go, que demonstra a arquitetura hexagonal na criação e consulta de produtos, com foco especial na implementação de validação de requisições utilizando *custom validators*. A aplicação é conteinerizada com Docker e Docker Compose, facilitando o ambiente de desenvolvimento e deploy.

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/Docker_Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker Compose">
  <img src="https://img.shields.io/badge/Architectural_Pattern-Hexagonal-blueviolet?style=for-the-badge" alt="Hexagonal Architecture">
  <img src="https://img.shields.io/badge/Validation-Custom-orange?style=for-the-badge" alt="Custom Validation">
</p>

## ✨ Visão Geral

Este projeto é um estudo prático da Arquitetura Hexagonal (Ports and Adapters) em Go, aplicado a um cenário de gerenciamento de produtos. A API permite a criação de novos produtos e a consulta de produtos existentes através de seu SKU. O diferencial e principal ponto de estudo desta aplicação reside na implementação de *custom validators* para garantir a integridade e conformidade dos dados recebidos nas requisições, demonstrando como criar regras de validação complexas e específicas para o domínio da aplicação.

## 🚀 Funcionalidades

* **Criação de Produto:** Endpoint para registrar novos produtos, incluindo nome, descrição, preço e SKU.
* **Consulta de Produto por SKU:** Endpoint para buscar detalhes de um produto específico utilizando seu SKU.
* **Validação Customizada de Requisições:** Implementação robusta de `custom validators` para validar o payload de entrada (e.g., formato de SKU, valores mínimos/máximos, etc.), garantindo que apenas dados válidos cheguem à camada de domínio.
* **Arquitetura Hexagonal:** Separação clara de responsabilidades entre as camadas (domínio, portas, adaptadores), promovendo alta coesão e baixo acoplamento.
* **Conteinerização:** Ambiente de desenvolvimento e produção facilmente configurável com Docker e Docker Compose.

## 🏗️ Arquitetura

A aplicação segue a Arquitetura Hexagonal, dividida nas seguintes camadas:

* **Domínio (Core/Business Logic):** Contém as entidades de negócio (Produto), interfaces de porta (`ProductServicePort`, `ProductRepositoryPort`) e a lógica de negócio principal.
* **Portas (Interfaces):** Define as interfaces que a camada de domínio espera de seus "drivers" (APIs, CLI) e "driven" (bancos de dados, serviços externos).
* **Adaptadores (Implementações):** Implementam as portas.
    * **Inbound Adapters (Drivers):** Ex: Adaptador HTTP (REST API) que consome as requisições e chama a porta de serviço.
    * **Outbound Adapters (Driven):** Ex: Adaptador de repositório (implementação do banco de dados) que persiste e recupera os dados.

Este design permite que a lógica de negócio permaneça independente de tecnologias externas (frameworks web, bancos de dados) e facilita a testabilidade.

## 🛠️ Tecnologias Utilizadas

* **Go (Golang)**
* **Docker**
* **Docker Compose**
* **Gin Gonic (ou similar, dependendo da sua escolha para a API)**
* **Go-playground/validator (para validação)**
* **SQLite (ou outro DB, se você usou algum)**

## 💻 Como Rodar

Certifique-se de ter o Docker e Docker Compose instalados em sua máquina.

1.  **Clone o repositório:**

    ```bash
    git clone [https://github.com/br4tech/go-custom-format.git](https://github.com/br4tech/go-custom-format.git)
    cd go-custom-format
    ```

2.  **Construa e Suba os Contêineres:**

    ```bash
    docker-compose up --build
    ```

    Isso construirá as imagens Docker (se necessário) e iniciará os contêineres para a aplicação Go e qualquer serviço de banco de dados configurado.

3.  **Acesse a Aplicação:**

    A API estará disponível em `http://localhost:8080` (ou a porta configurada no seu `docker-compose.yml`).

## 🎯 Custom Validation em Ação

O ponto chave deste projeto é a demonstração de como estender e usar o pacote de validação (ex: `go-playground/validator`) para criar regras de validação que vão além das tags padrão (`required`, `min`, `max`, etc.).

### Exemplo de Uso (com Validação)

#### Criar Produto

`POST /products`

```json
{
    "name": "Smartphone X",
    "description": "O mais novo smartphone com câmera de 108MP.",
    "price": 1299.99,
    "sku": "SPH-X-001"
}

# Validação de SKU (Exemplo de Custom Validation)

Pode ser configurado um validador customizado para garantir que o SKU siga um padrão específico (ex: `[A-Z]{3}-[A-Z]{1}-[0-9]{3}`).

Se o SKU for inválido, a API retornará um erro **400 Bad Request** com uma mensagem descritiva.

## Consultar Produto por SKU

**GET** `/products/{sku}`

**Exemplo:**  
`GET /products/SPH-X-001`

### Resposta de Sucesso

```json
{
  "id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
  "name": "Smartphone X",
  "description": "O mais novo smartphone com câmera de 108MP.",
  "price": 1299.99,
  "sku": "SPH-X-001",
  "created_at": "2023-10-27T10:00:00Z"
}

Resposta de Erro (Produto Não Encontrado)

```json
{
  "message": "Product with SKU SPH-X-002 not found."
}


## 📚 Como o Custom Validator é Implementado

No código, você encontrará:

### 1. Definição do struct de requisição

Onde as tags de validação padrão e customizadas são aplicadas.

```go
type CreateProductRequest struct {
    Name        string  `json:"name" validate:"required,min=3,max=100"`
    Description string  `json:"description" validate:"required,min=10,max=500"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    SKU         string  `json:"sku" validate:"required,skuformat"` // <<--- Aqui está a validação customizada
}

### 2. Função de validação customizada
Uma função que implementa a lógica da sua regra de validação (skuformat, por exemplo).

```go
// Exemplo simplificado de custom validator (pode estar em um pacote de utilitários ou de validação)
func validateSKUFormat(fl validator.FieldLevel) bool {
    sku := fl.Field().String()
    // Implemente sua lógica de regex ou qualquer outra aqui
    match, _ := regexp.MatchString("^[A-Z]{3}-[A-Z]{1}-[0-9]{3}$", sku)
    return match
}

### 3. Registro do validador customizado
Onde você registra sua função de validação com a instância do go-playground/validator.

```go
// Exemplo de como registrar (geralmente feito na inicialização da aplicação)
validate := validator.New()
validate.RegisterValidation("skuformat", validateSKUFormat)

4. Uso do validador no handler
Antes de processar a requisição, o validate.Struct() é chamado para aplicar as validações.