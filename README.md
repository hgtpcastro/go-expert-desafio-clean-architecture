# go-expert-desafio-clean-architecture
Pós Go Expert desafio Clean Architecture

## Subir os serviços de banco de dados MySql e RabbitMq
```
make docker-up
```

## Rodar as migrations
```
make migrate-up
```

## Subir os servidores Http Grpc e GraphQl
```
make run-app
```
## Consumindo a API

### 1 - Via Http (http://localhost:8000/order)
```
cd ./api/create_order.http
```
```
cd ./api/get_orders.http
```

### 2 - Via Grpc (bash Evans)
```
evans -r repl
```
```
call CreateOrder
```
```
call GetOrders
```

### 3 - Via GraphQl (http://localhost:8080)

#### 3.1 - Mutation
```
mutation createOrder{
  createOrder(input: {id: "c", Price: 300.0, Tax: 2.0}) {
    id
    Price
    Tax
    FinalPrice
  } 
}
```
```
query getOrders{
  getOrders {
    id
    Price
    Tax
    FinalPrice
  } 
}
```
## Tools

### Wire (na pasta do arquivo wire.go)
```
wire 
```

### Grpc
```
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

### GraphQl
```
go run github.com/99designs/gqlgen generate
```

