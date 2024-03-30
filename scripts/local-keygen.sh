#!/bin/bash 
curl --location --request POST 'http://127.0.0.1:7044/validator/params' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String1",
	"Moniker": "Moniker String1",
	"Key": "3141592653589793238462643383279502884197169399375105820974944591",
	"ip": "http://127.0.0.1:7045",
	"partyPassword": "password string",
	"Threshold": 2
}'

curl --location --request POST 'http://127.0.0.1:7044/validator/params' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String2",
	"Moniker": "Moniker String2",
	"Key": "3141592653589793238462643383279502884197169399375105820974944592",
	"ip": "http://127.0.0.1:7046",
	"partyPassword": "password string",
	"Threshold": 2
}'

curl --location --request POST 'http://127.0.0.1:7044/validator/params' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String3",
	"Moniker": "Moniker String3",
	"Key": "3141592653589793238462643383279502884197169399375105820974944593",
	"ip": "http://127.0.0.1:7047",
	"partyPassword": "password string",
	"Threshold": 2
}'


curl --location --request POST 'http://127.0.0.1:7044/validator/start' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String",
	"Moniker": "Moniker String",
	"Key": "314159265358979323846264338327950288419716939937510582097494459",
	"IP": "IP URL for the Genesis server",
	"partyPassword": "password string",
	"Threshold": 3
}'
