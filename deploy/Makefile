migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/test?sslmode=disable -path ../db/migrations down

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/test?sslmode=disable -path ../db/migrations up

#migrate_create:
#	migrate create -ext sql -dir ../db/migrations -seq init_db

BINARY=serviceApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: service
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

service:
	@echo "Building service binary..."
	cd ./../. && env GOOS=linux CGO_ENABLED=0 go build -o ${BINARY} ./cmd/app
	@echo "Done!"
