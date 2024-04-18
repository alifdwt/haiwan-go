server:
	go run cmd/app/main.go

swag:
	swag init -g internal/app/app.go

mock:
	mockgen -package mockdb -source internal/service/interfaces.go -destination internal/mock/service.go ...

fe-admin:
	cd web/admin && npm install && npm run dev

fe-shop:
	cd web/shop && npm install && npm run dev

.PHONY: server swag mock fe-admin