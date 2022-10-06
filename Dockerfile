FROM golang:latest
ENV CAP, CURRENCY
RUN mkdir /funmoney && cd /funmoney
WORKDIR /funmoney
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY * ./
RUN go build .
EXPOSE 8080
CMD ["./funmoney"]