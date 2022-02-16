#!/bin/bash

cd "$(dirname $0)" || exit 1

openssl genrsa \
  -out ../certs/server.key \
  2048

openssl req \
  -new \
  -subj "/C=CN/ST=Tianjin/L=Tianjin/O=ZhangYuheng Inc/CN=Test Server" \
  -key ../certs/server.key \
  -out ../certs/server.csr

openssl x509 \
    -req \
    -days 3650 \
    -CA ../certs/ca.crt \
    -CAkey ../certs/ca.key \
    -set_serial 01 \
    -extfile server_v3_ext.conf \
    -in ../certs/server.csr \
    -out ../certs/server.crt