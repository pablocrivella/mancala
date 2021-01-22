DATABASE_URL="postgres://mancala-app:p4ssw0rd@0.0.0.0:5432/mancala-db?sslmode=disable
API_SPEC_PATH="./api/openapi-v2/spec.yml"

dev.up:
	docker-compose up --remove-orphans

server:
	env DATABASE_URL=$(DATABASE_URL) go run ./cmd/api --schema http

test:
	go test ./...

db.reset: db.drop db.up

db.drop:
	env DATABASE_URL=$(DATABASE_URL) dbmate drop

db.up:
	env DATABASE_URL=$(DATABASE_URL) dbmate up

db.gen:
	sqlc generate

api.gen:
	rm -rf api/openapi-v2/models
	rm -rf api/openapi-v2/restapi
	rm -rf api/openapi-v2/cmd
	swagger generate server -q --target ./api/openapi-v2 --name mancala --spec $(API_SPEC_PATH) --flag-strategy=flag --exclude-main

api.validate:
	swagger validate $(API_SPEC_PATH)
