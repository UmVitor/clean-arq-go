# Clean Architecture Order System

Este projeto implementa um sistema de pedidos (orders) utilizando clean arquitecture com múltiplas interfaces de comunicação:
- REST API
- gRPC
- GraphQL

## Pré-requisitos

- Go 1.16+
- MySQL
- RabbitMQ

## Configuração

1. Clone o repositório
2. Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
```

3. Ajuste as configurações conforme necessário para seu ambiente

## Executando a aplicação

```bash
cd cmd/ordersystem
go run main.go wire_gen.go
```

A aplicação iniciará três servidores:
- REST API: http://localhost:8000
- gRPC: localhost:50051
- GraphQL: http://localhost:8080

## Endpoints disponíveis

### REST API

#### Criar um pedido (Create Order)

```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{"id":"1", "price": 100.0, "tax": 10.0}'
```

#### Listar pedidos (List Orders)

```bash
curl -X GET http://localhost:8000/order
```

### gRPC

Para testar o gRPC, você pode usar ferramentas como [grpcurl](https://github.com/fullstorydev/grpcurl) ou [BloomRPC](https://github.com/uw-labs/bloomrpc).

#### Listar pedidos (List Orders)

```bash
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```

### GraphQL

O GraphQL Playground está disponível em: http://localhost:8080

#### Criar um pedido (Create Order)

```graphql
mutation {
  createOrder(input: {id: "1", Price: 100.0, Tax: 10.0}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Listar pedidos (List Orders)

```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## Estrutura do projeto

O projeto segue os princípios da Clean Architecture:

- `cmd/ordersystem`: Ponto de entrada da aplicação
- `configs`: Configurações da aplicação
- `internal/entity`: Entidades de domínio
- `internal/usecase`: Casos de uso da aplicação
- `internal/infra`: Implementações de infraestrutura
  - `database`: Repositórios
  - `web`: Handlers HTTP
  - `grpc`: Serviços gRPC
  - `graph`: Resolvers GraphQL

## Portas dos serviços

- REST API: 8000 (configurável via WEB_SERVER_PORT)
- gRPC: 50051 (configurável via GRPC_SERVER_PORT)
- GraphQL: 8080 (configurável via GRAPHQL_SERVER_PORT)
