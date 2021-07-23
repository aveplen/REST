.PHONY: build
.PHONY: run
.PHONY: exe

run:
	go build -v ./cmd/
	./cmd.exe

build:
	go build -v ./cmd/