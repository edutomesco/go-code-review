up:
	docker compose -f ./infrastructure/local/docker-compose.yaml up -d

down:
	docker compose -f ./infrastructure/local/docker-compose.yaml down

mocks:
	mockgen -source=internal/interfaces/coupon_repository.go -destination internal/interfaces/mocks/mock_coupon_repository.go

cache:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...

test:
	go test ./...
