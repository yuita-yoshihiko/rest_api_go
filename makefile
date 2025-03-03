RUN=docker compose run --service-ports --rm --workdir="/rest-api-go" rest-api-go

migrate/new:
	${RUN} sh -c "sql-migrate new ${FILE_NAME}"

migration/status:
	${RUN} go tool sql-migrate status --env='development'

migration/up:
	${RUN} go tool sql-migrate up --env='development'

migration/down:
	${RUN} go tool sql-migrate down --env='development'

sqlc:
	${RUN} go tool sqlc generate

psql:
	psql -h 127.0.0.1 -p 5732 -U user rest-api-go-db

test-db-up:
	docker compose -f ./docker-compose.test-db.yml up -d
	sleep 5
	${RUN} sh -c "sql-migrate up --env='test'"

test-db-down:
	docker compose -f ./docker-compose.test-db.yml down

output-mg:
	${RUN} sh -c "go run internal/skeleton/migration/main.go"

output-sc:
	${RUN} sh -c "go run internal/skeleton/schema/main.go"

output-uc:
	${RUN} sh -c "go run internal/skeleton/usecase/main.go"

output-cv:
	${RUN} sh -c "go run internal/skeleton/converter/main.go"

output-rp:
	${RUN} sh -c "go run internal/skeleton/repository/main.go"

output-db:
	${RUN} sh -c "go run internal/skeleton/database/main.go"

output-all-skeleton:
	make output-mg
	make output-sc
	make output-uc
	make output-cv
	make output-rp
	make output-db