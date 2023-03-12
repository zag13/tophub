APP = UNSPECIFIED

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: api
# generate api proto
api:
	@if [ "$(APP)" == "UNSPECIFIED" ]; then \
		find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'; \
	else \
  		cd app/$(APP)/service && pwd && $(MAKE) api; \
  	fi

.PHONY: config
# generate internal proto
config:
	@if [ "$(APP)" == "UNSPECIFIED" ]; then \
		find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) config'; \
	else \
  		cd app/$(APP)/service && pwd && $(MAKE) config; \
  	fi

.PHONY: wire
# generate wire
wire:
	@if [ "$(APP)" == "UNSPECIFIED" ]; then \
		find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'; \
	else \
  		cd app/$(APP)/service && pwd && $(MAKE) wire; \
  	fi