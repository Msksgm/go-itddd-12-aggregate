up:
	docker compose up -d
down:
	docker compose down --remove-orphans
destroy:
	docker compose down --rmi all --volumes --remove-orphans
create-migration:
	docker compose exec app ./migrate create -ext sql -dir migrations -seq ${file}
run-migration:
	docker compose exec app bash db-migration.sh
run:
	docker compose exec app go run main.go -usecase=${usecase}
test:
	docker compose exec app go test ./...
