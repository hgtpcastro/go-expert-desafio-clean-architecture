.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: migrate-up
migrate-up:	
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up 

.PHONY: migrate-down
migrate-down:	
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down 	

.PHONY: run-app
run-app:
	go run "./cmd/ordersystem/main.go" "./cmd/ordersystem/wire_gen.go"

.PHONY: protoc
run-protoc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto