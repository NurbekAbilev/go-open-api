include .env
export $(shell sed 's/=.*//' .env)

install_migrator:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

# Called as:
# $ migrate_create name=create_table_users
migrate_create:
	migrate create -ext sql -dir migrations $(name)

migrate_up:
	migrate -path migrations -database $(DB_URL) up

migrate_down:
	migrate -path migrations -database $(DB_URL) down
