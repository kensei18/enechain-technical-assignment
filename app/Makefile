generate_web:
	gqlgen generate --config ./gqlgen.web.yml

generate_admin:
	gqlgen generate --config ./gqlgen.admin.yml

generate_migration:
	migrate create -ext sql -dir migration $(name)

migrate:
	migrate -database postgres://postgres:password@localhost:5432/app?sslmode=disable -path migration up
	migrate -database postgres://postgres:password@localhost:5432/app_test?sslmode=disable -path migration up

rollback:
	migrate -database postgres://postgres:password@localhost:5432/app?sslmode=disable -path migration down 1
	migrate -database postgres://postgres:password@localhost:5432/app_test?sslmode=disable -path migration down 1

boil:
	sqlboiler --add-soft-deletes psql
