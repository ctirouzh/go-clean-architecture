.PHONY: server
server:
	go run cmd/main.go

.PHONY: test
test:
	go test ./... -cover -race
