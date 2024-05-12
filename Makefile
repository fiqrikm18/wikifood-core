dev:
	swag init -g cmd/app.go && go run main.go

swag-fmt:
	swag fmt