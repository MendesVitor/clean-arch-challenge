# Go Orders

## Descrição

Este projeto implementa um sistema de listagem de pedidos utilizando três tipos de transporte: REST, gRPC e GraphQL.

## Configuração do Ambiente

1. **Instale o Docker e Docker Compose**

2. **Clone o repositório**

    ```sh
    git clone <repo_url>
    cd go-orders
    ```

3. **Suba os serviços com Docker Compose**

    ```sh
    docker-compose up
    ```

4. **Acesse os serviços**
    - REST: `http://localhost:8080/order`
    - gRPC: `localhost:50051`
    - GraphQL: `http://localhost:8081/query`

## Endpoints

### REST

-   **Listar Pedidos**
    ```http
    GET /order
    ```

### gRPC

-   **Listar Pedidos**
    ```protobuf
    service OrderService {
    rpc ListOrders (Empty) returns (OrderList) {}
    }
    ```

### GraphQL

-   **Listar Pedidos**
    ```graphql
    query {
        listOrders {
            id
            description
            price
            created_at
        }
    }
    ```
