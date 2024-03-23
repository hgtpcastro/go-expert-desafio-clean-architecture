# go-expert-desafio-clean-architecture
Pós Go Expert desafio Clean Architecture

## Subir os serviços de banco de dados MySql e RabbitMq
```
docker-compose up -d 
```

## Subir os servidores Http Grpc e GraphQl
```
cd ./cmd/ordersystem/
```
```
go run main.go wire_gen.go
```

## Preparar o banco de dados (docker)

### Pega o nomes do container MySql
```
docker container ls
```

### Entra no container MySql
```
docker exec -it mysql bash
```
### Bash MySql
```
bash-4.2# mysql -uroot -p (senha root)
mysql> show databases;
mysql> use orders;
mysql> show tables;
```

### Cria a tabela (bash MySql)
```
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));
```
``` 
show tables;
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

