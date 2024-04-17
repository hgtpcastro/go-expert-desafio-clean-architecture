# go-expert-desafio-clean-architecture
Pós Go Expert desafio Clean Architecture

## 1 - Subir os serviços de banco de dados MySql e RabbitMq
```
make docker-up
```

## 2 - Rodar as migrations
```
make migrate-up
```

## 3 - Subir os servidores
```
make run-app
```
### Portas disponíveis para cada servidor
- Web Server: **8000**
- Grpc Server: **50051**
- GraphQl Server: **8080**

## Consumindo os servidores

### Web Server
#### Utilizar a extensão REST Client
```
cd ./api/create_order.http
```
```
cd ./api/get_orders.http
```

### Grpc server
#### Utilizar o Evans para as requisições
```
evans -r repl
```
```
call CreateOrder
```
```
call GetOrders
```

### GraphQl Server
#### Utilizar Playground do GraphQl

### Mutation
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
### Query
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