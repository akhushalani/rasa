postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root rasa

dropdb:
	docker exec -t postgres17 dropdb rasa

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rasa?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rasa?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc