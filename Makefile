E = UNKNOWN
S = UNKNOWN

.PHONY: cli-build
cli-build:
	@echo "building cli"
	cd cli && go build -o ../bin/topcli ./*.go;

.PHONY: cli-run
cli-run:
	@echo "running cli"
	cd cli && go run main.go spider;

.PHONY: client-build
client-build:
	@echo "building client"

.PHONY: client-run
client-run:
	@echo "running client"

.PHONY: server-gen
server-gen:
	@echo "gen server dal..."
	cd server && go run cmd/gen/*.go

.PHONY: server-build
server-build:
	@echo "building server"

.PHONY: server-run
server-run:
	@echo "running server"
	cd server && go run *.go

.PHONY: tidy
tidy:
	@echo "tidy"
	cd cli && go mod tidy; \
	cd ../server && go mod tidy;
