#!/bin/bash 

# Servers Info
# 1. Genesis-Server - 18.229.137.38:7044
# 2. Non-Genesis-Server-1 - 18.228.192.203:7044
# 3. Non-Genesis-Server-2 - 54.233.170.242:7044
# 4. Non-Genesis-Server-3 - 18.231.112.8:7044

# Sending Keysign request to Genesis-Server
curl --location --request POST 'http://18.229.137.38:7044/validator/sign' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "1022"
}'
echo 