.PHONY: build-client
build-client:
	@echo "building client"
	cd client && pnpm build

.PHONY: run-client
run-client:
	@echo "running client"
	cd client && pnpm dev

.PHONY: build-cli
build-cli:
	@echo "building cli"
	cd cli && go build -o ../bin/topcli ./*.go;

.PHONY: run-cli
run-cli:
	@echo "running cli"
	cd cli && go run main.go spider;

.PHONY: gen-server
gen-server:
	@echo "generate server dal..."
	cd server && go run cmd/gen/*.go

.PHONY: run-server
run-server:
	@echo "running server"
	cd server && go run *.go

.PHONY: tidy
tidy:
	@echo "tidy"
	cd cli && go mod tidy; \
	cd ../server && go mod tidy;

.PHONY: dkc-up
dkc-up:
	@echo "docker compose up..."
	cd deploy/docker-compose && docker compose up --build