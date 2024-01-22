run:
	go run ./cmd/web
migrate:
	psql -U postgres -c "CREATE DATABASE storage;"
	psql -U postgres -d storage -a -f migrations/001_init.up.sql