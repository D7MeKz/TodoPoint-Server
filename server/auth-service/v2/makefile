hello:
	echo "auth-service/v2"

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/main main.go

clean:
	rm -rf ./bin

test:
	go test ./...
