db:
	docker run --name tuvebes_db -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db -d postgres

migrate:
	migrate -path ./migrations -database 'postgres://username:password@localhost:5432/db?sslmode=disable' up

run:
	go run cmd/main.go

