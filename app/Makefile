generate_web:
	gqlgen generate --config ./gqlgen.web.yml

generate_admin:
	gqlgen generate --config ./gqlgen.admin.yml

generate_migration:
	migrate create -ext sql -dir migration $(name)

migrate:
	migrate -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/app?sslmode=disable -path migration up
	migrate -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/app_test?sslmode=disable -path migration up

rollback:
	migrate -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/app?sslmode=disable -path migration down 1
	migrate -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/app_test?sslmode=disable -path migration down 1
