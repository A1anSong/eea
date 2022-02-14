all: server client

server: $(shell find server -type f)
	cd server; go build -o ../bin/server
api:

test:
	cd server; go test -v
web:

dev: server web
	cp config/server.yaml bin/server.yaml
	cd bin; ./server

.PHONY: server
