include .env

.PHONY: test-env migrate-up migrate-down migrate-create

# Command to verify env variables are loaded
test-env:
	@echo "Database URL: $(DATABASE_URL)"


migrate-up:
	migrate -path db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DATABASE_URL)" down

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-force:
	migrate -path db/migrations -database "$(DATABASE_URL)" force $(version)

migrate-version:
	migrate -path db/migrations -database "$(DATABASE_URL)" version

migrate-down-steps:
	migrate -path db/migrations -database "$(DATABASE_URL)" down $(steps)
