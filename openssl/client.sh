#!/bin/bash

cd "$(dirname $0)" || exit 1

openssl genrsa \
  -out ../certs/client.key \
  2048

openssl req \
  -new \
  -subj "/C=CN/ST=Tianjin/L=Tianjin/O=ZhangYuheng Inc/CN=Test Client" \
  -key ../certs/client.key \
  -out ../certs/client.csr

openssl x509 \
    -req \
    -days 3650 \
    -CA ../certs/ca.crt \
    -CAkey ../certs/ca.key \
    -set_serial 01 \
    -extfile client_v3_ext.conf \
    -in ../certs/client.csr \
    -out ../certs/client.crt