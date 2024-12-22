run:
	@go run cmd/main/main.go

up:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/simbir_go?sslmode=disable' up

down:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/simbir_go?sslmode=disable' down