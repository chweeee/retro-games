GIT_ROOT = $(shell git rev-parse --show-toplevel)

hello:
	@echo $(GIT_ROOT)

build:
	go build -o snake $(GIT_ROOT)/cmd/snake/main.go

run:
	go run $(GIT_ROOT)/cmd/snake/main.go

test:


