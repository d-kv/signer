FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go build -o query_handler_main ./query-handler/cmd/app/main.go

EXPOSE 8081

CMD ["./query_handler_main"]
