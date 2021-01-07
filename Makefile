bin/consumer:
	GOOS=darwin GOARCH=amd64 go build -o bin/consumer ./consumer/main.go
	chmod +x bin/consumer

bin/producer:
	GOOS=darwin GOARCH=amd64 go build -o bin/producer ./producer/main.go
	chmod +x bin/producer

clean:
	rm -rf bin

.PHONY: clean