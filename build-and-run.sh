#!/usr/bin/env sh
buildDocker() {
    docker build -t funmoney-frontend frontend && docker build -t funmoney .
}

runDockerCompose() {
    docker-compose up --remove-orphans -d
}

buildDocker && runDockerCompose