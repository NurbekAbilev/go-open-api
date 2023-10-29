include .env
export $(shell sed 's/=.*//' .env)

#  migrate -database postgres://postgres:example@localhost:5432/postgres?sslmode=disable -path
install_migrator:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

migrate_create:
	migrate create -ext sql -dir migrations $(name)

migrate_up:
	migrate -path migrations -database $(DB_URL) up
# migrate_up:
	
