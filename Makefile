.PHONY: run

run:
	@go run $(wildcard *.go)

.PHONY: clean
clean:
	@rm -f $(wildcard *.out)