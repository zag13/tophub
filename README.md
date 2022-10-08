# tophub

## Create a service

```
# Create a template project
kratos new app/{{serverName}} --nomod

# Add a proto template
kratos proto add api/{{serverName}}/v1/{{serverName}}.proto

# Generate the proto code
kratos proto client api/{{serverName}}/v1/{{serverName}}.proto

# Generate the source code of service by proto file
kratos proto server api/{{serverName}}/v1/{{serverName}}.proto -t app/{{serverName}}/internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
