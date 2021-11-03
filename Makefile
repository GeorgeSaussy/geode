test:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

format:
	go fmt

clean:
	go clean -testcache
