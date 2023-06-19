`docker build -t pehgomess/goapp-hello .`

`docker run -p 9000:9000 pehgomess/goapp-hello`

`docker build -t pehgomess/nginx-mtls .`

rodar os comandos de certificado dentro do diretorio nginx/certs

```bash

openssl genrsa -des3 -out ca.key 4096

openssl rsa -in ca.key -out ca.key
openssl req -new -x509 -days 3650 -key ca.key -subj "/CN=*.skypoc" -out ca.crt
 
printf test > passphrase.txt
openssl genrsa -des3 -passout file:passphrase.txt -out server.key 2048
openssl req -new -passin file:passphrase.txt -key server.key -subj "/CN=*.skypoc" -out server.csr
 
openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

#CN=IP-CLIENT
printf test > client_passphrase.txt
openssl genrsa -des3 -passout file:client_passphrase.txt -out client.key 2048
openssl rsa -passin file:client_passphrase.txt -in client.key -out client.key
openssl req -new -key client.key -subj "/CN=IP-CLIENT" -out client.csr
 
openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out client.crt

```
`docker-compose up`

Adicionar no hosts a entrada do localhost com o nome do dominio correspondente do dominio do CN do wildcard 

`curl --key client.key --cert client.crt https://nginx.skypoc:8443/teste/v1/api/mtls --insecure`