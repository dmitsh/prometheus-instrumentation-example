all: server client

server:
	go build -o demo-srv ./cmd/server

client:
	go build -o demo-cln ./cmd/client

clean:
	rm -f ./demo-srv ./demo-cln

.PHONY: server client clean
