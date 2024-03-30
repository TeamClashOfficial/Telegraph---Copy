#!/bin/bash 

# Servers Info
# 1. Genesis-Server - 18.229.137.38:7044
# 2. Non-Genesis-Server-1 - 18.228.192.203:7044
# 3. Non-Genesis-Server-2 - 54.233.170.242:7044
# 4. Non-Genesis-Server-3 - 18.231.112.8:7044


# Sending Non-Genesis-Server-1 info to Genesis-Server
curl --location --request POST 'http://18.229.137.38:7044/validator/params' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String1",
	"Moniker": "Moniker String1",
	"Key": "3141592653589793238462643383279502884197169399375105820974944591",
	"ip": "http://18.228.192.203:7044",
	"partyPassword": "password string",
	"Threshold": 2
}'

# Sending Non-Genesis-Server-2 info to Genesis-Server
curl --location --request POST 'http://18.229.137.38:7044/validator/params' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String2",
	"Moniker": "Moniker String2",
	"Key": "3141592653589793238462643383279502884197169399375105820974944592",
	"ip": "http://54.233.170.242:7044",
	"partyPassword": "password string",
	"Threshold": 2
}'

# Sending Non-Genesis-Server-3 info to Genesis-Server
curl --location --request POST 'http://18.229.137.38:7044/validator/params' \
--header 'Content-Type: application/json' \
--data-raw '{
	"isGenesis": true,
	"ID": "ID String3",
	"Moniker": "Moniker String3",
	"Key": "3141592653589793238462643383279502884197169399375105820974944593",
	"ip": "http://18.231.112.8:7044",
	"partyPassword": "password string",
	"Threshold": 2
}'

# Sending keygen request to Genesis-Server
curl --location --request POST 'http://18.229.137.38:7044/validator/start' \
--header 'Content-Type: application/json' \
--data-raw '{
	"partyPassword": "password string"
}'

echo