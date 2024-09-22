DB_URL=postgresql://root:secret@localhost:5432/numerisbookdb?sslmode=disable

postgres:
	docker compose up

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:	
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

start:
	docker compose start

stop:
	docker compose stop

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres new_migration migrateup migratedown migrateup1 migratedown1 start stop sqlc test