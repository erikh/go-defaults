all: auto-test

auto-test:
	reflex -r '\.go$$' -- make test

test:
	go test -v ./...

.PHONY: test auto-test all
