GOLANGCI_LINT_VERSION = v1.53.3

.PHONY: buf
buf:
	@sh ./scripts/buf.sh
	@buf lint proto
	@buf format proto -w
	@buf generate proto

.PHONY: test
test:
	go test ./... -v --bench . --benchmem --coverprofile=cover.out

testsimple:
	go test ./... --bench . --benchmem --coverprofile=cover.out

testrace:
	go test -race ./... -v --bench .

testcovhttp:
	go test ./... -v --coverprofile=cover.out && go tool cover -html=cover.out

lint:
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} golangci-lint run -v

check: lint test

run:
	go run $$(pwd)/cmd/todo/main.go
runwiretap:
	go run $$(pwd)/cmd/wiretap/main.go

up:
	docker-compose -f $$(pwd)/build/package/docker-compose.yml up -d
	@ echo "view jaeger at http://localhost:16686"

down:
	docker-compose -f $$(pwd)/build/package/docker-compose.yml down

downup: down up

dbup:
	docker-compose -f $$(pwd)/build/package/docker-compose.yml up -d db

jaegerup:
	docker-compose -f $$(pwd)/build/package/docker-compose.yml up -d jaeger
	@ echo "view jaeger at http://localhost:16686"

natsup:
	docker-compose -f $$(pwd)/build/package/docker-compose.yml up -d nats
