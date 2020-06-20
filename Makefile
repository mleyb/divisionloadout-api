build:
	go build -o bin/service main.go

format:
	go fmt ./...

test:
	go test ./...

run:
	go run main.go

zip:
	zip function.zip service