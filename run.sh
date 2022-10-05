#!/bin/env sh
docker build -t funmoney . && docker-compose up --remove-orphans