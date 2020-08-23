GIT_ROOT = $(shell git rev-parse --show-toplevel)

hello:
	@echo $(GIT_ROOT)

build:
	go build -o snake $(GIT_ROOT)/cmd/snake/main.go
	go build -o pong $(GIT_ROOT)/cmd/pong/main.go

snake:
	go run $(GIT_ROOT)/cmd/snake/main.go

pong:
	go run $(GIT_ROOT)/cmd/pong/main.go

