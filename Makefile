server:
	go run cmd/app/main.go

swag:
	swag init -g internal/app/app.go

mock:
	mockgen -package mockdb -source internal/service/interfaces.go -destination internal/mock/service.go ...

.PHONY: server swag mock