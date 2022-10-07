#!/usr/bin/env sh
buildDocker() {
    docker build -t funmoney-frontend frontend && docker build -t funmoney .
}

runDockerCompose() {
    ENV=$ENV docker-compose up --remove-orphans -d
    docker-compose restart funmoney
}

buildDocker && runDockerCompose