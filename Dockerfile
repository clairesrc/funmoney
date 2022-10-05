FROM golang:latest
ENV CAP, CURRENCY
RUN mkdir /funmoney && cd /funmoney
WORKDIR /funmoney
COPY * /funmoney/
RUN go get .
RUN go build .
EXPOSE 8080
CMD ["./funmoney"]