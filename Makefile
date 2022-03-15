all: server client

server: $(shell find server -type f)
	cd server; go build -o ../bin/server
api:

test:
	cd server; go test -v
web:

dev_server:
	cd server && air
dev_web:
	cd web && yarn dev
dev: dev_server dev_web

.PHONY: server
