#!/bin/env sh

POST_BODY=<<EOF
{
    "comment": "Automated monthly disbursement", 
    value: -$CAP
}
EOF;

wget -q -O- --post-data $POST_BODY -H "http://localhost:8080/" --header "content-type: application/json"
