.PHONY: run

run:
	@go run $(wildcard *.go)

.PHONY: clean
clean:
	@rm -f $(wildcard *.out)

.PHONY: test
test:
	go test ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: test-coverage-html
test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	explorer.exe coverage.html