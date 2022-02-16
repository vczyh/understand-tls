#!/bin/bash

cd "$(dirname $0)" || exit 1


openssl genrsa \
  -out ../certs/ca.key \
  2048

openssl req \
  -x509 \
  -new \
  -days 3650 \
  -subj "/C=CN/ST=Tianjin/L=Tianjin/O=ZhangYuheng Inc/CN=ZhangYuheng Root CA" \
  -key ../certs/ca.key \
  -out ../certs/ca.crt