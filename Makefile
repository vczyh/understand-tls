.PHONY: pre openssl-ca openssl-server openssl-client openssl

all: build

pre:
	chmod -R +x openssl

openssl-ca: pre
	openssl/ca.sh

openssl-server: pre
	openssl/server.sh

openssl-client: pre
	openssl/client.sh

openssl: openssl-ca openssl-server openssl-client
