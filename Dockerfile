FROM golang:1.21.5

WORKDIR /go/src/httpserver

COPY . .
RUN go mod download

RUN go build -o httpserver .

EXPOSE 8888

CMD ["./httpserver"]