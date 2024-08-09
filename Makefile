start:
	docker-compose up --build -d postgres
migrateUp:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
migrateStatus:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" status
down:
	docker-compose down