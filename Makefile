E = UNKNOWN
S = UNKNOWN

.PHONY: cli-build
# building cli
cli-build:
	@echo "building cli"
	cd cli && pwd && go build -o ../bin/ ./...;

.PHONY: cli-run
# running cli
cli-run:
	@echo "running cli"
	@if [ "$(E)" == "UNKNOWN" && "$(S)" == "UNKNOWN" ]; then \
		cd cli && pwd && go run main.go spider zhihu file; \
	else \
		cd cli && pwd && go run main.go spider $(E) $(S); \
	fi

.PHONY: client-build
# building client
client-build:
	@echo "building client"

.PHONY: client-run
# running client
client-run:
	@echo "running client"

.PHONY: server-build
# building server
server-build:
	@echo "building server"

.PHONY: server-run
# running server
server-run:
	@echo "running server"