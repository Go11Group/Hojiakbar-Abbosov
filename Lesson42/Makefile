DB_URL = postgres://postgres:root@localhost:5432/nil

run:
	go run main.go
migrate_file:
	migrate create -ext sql -dir migrations -seq nil
migrate_up:
	migrate -path migrations -database ${DB_URL} up
migrate_down:
	migrate -path migrations -database ${DB_URL} down
migrate_force:
	migrate -path migrations -database ${DB_URL} force 1
c:
	clear