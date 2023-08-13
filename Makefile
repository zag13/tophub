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