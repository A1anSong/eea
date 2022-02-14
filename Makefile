all: server client

server: $(shell find server -type f)
	cd server; go build -o ../bin/server
api:

test:
	cd server; go test -v
web:

dev: server web
	cp configs/server_dev.yaml bin/server.yaml
	./bin/server

.PHONY: server
