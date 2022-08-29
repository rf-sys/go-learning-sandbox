migrate-up:
	migrate -path db/migration -database "postgresql://colbeypq:44wKzGteZg8ZfuVTzZ@localhost:5432/vinyl-db?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://colbeypq:44wKzGteZg8ZfuVTzZ@localhost:5432/vinyl-db?sslmode=disable" -verbose down -all
