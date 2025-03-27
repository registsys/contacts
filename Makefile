build:
	CGO_ENABLED=0 go build -o bin/contacts ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v ./...