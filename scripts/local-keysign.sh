#!/bin/bash 
curl --location --request POST 'http://127.0.0.1:7044/validator/sign' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "1022"
}'
