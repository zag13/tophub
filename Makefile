.PHONY: run-client
run-client:
	@echo "running client"
	cd client && pnpm dev

.PHONY: gen-cli
gen-cli:
	@echo "generate cli dal..."
	cd cli && go run cmd/gen/*.go

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

.PHONY: protoc
protoc:
	@echo "protoc ..."
	cd server && protoc --proto_path=./api/protos --proto_path=./third_party \
		--go_out=paths=source_relative:./api/types \
		--openapi_out=default_response=true:. \
		./api/protos/*.proto && \
	protoc-go-inject-tag -input=./api/types/*.pb.go

.PHONY: dkc-up
dkc-up:
	@echo "docker compose up..."
	cd deploy/docker-compose && docker compose up --build -d