FROM golang:latest
ENV CAP, CURRENCY
RUN mkdir /funmoney && cd /funmoney
WORKDIR /funmoney
COPY * /funmoney/
RUN go build .
CMD ["./funmoney"]