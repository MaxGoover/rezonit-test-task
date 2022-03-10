create-db:
	docker exec -it postgres14 createdb --username=root --owner=root rezonit_test_task

drop-db:
	docker exec -it postgres14 dropdb rezonit_test_task

migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rezonit_test_task?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rezonit_test_task?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go