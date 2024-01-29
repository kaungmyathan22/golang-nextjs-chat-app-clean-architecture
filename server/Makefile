
run:
	go run ./cmd/main.go
migrateup:
	migrate -path db/migrations -database "postgresql://admin:admin@localhost:5432/go-chat?sslmode=disable" -verbose up