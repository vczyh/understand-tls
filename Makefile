.PHONY: pre clean openssl-ca openssl-server openssl-client openssl go server client


pre:
	chmod -R +x openssl

clean:
	rm -rf certs/*

openssl-ca: pre
	openssl/ca.sh

openssl-server: pre
	openssl/server.sh

openssl-client: pre
	openssl/client.sh

openssl: openssl-ca openssl-server openssl-client

go:
	go run ./gotls

server:
	go run ./server

client:
	go run ./client
