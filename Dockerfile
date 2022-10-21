FROM golang:latest
RUN apt-get update && apt-get install -y cron
ENV CAP, CURRENCY, MONGODB_CONNECTION_URI
RUN mkdir /funmoney && cd /funmoney
WORKDIR /funmoney
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY * ./
RUN go build .
COPY crontab /var/spool/cron/crontabs/root
RUN cron -l 8
EXPOSE 8080
CMD ["./funmoney"]