FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go build -o api_main ./api/cmd/app/main.go

EXPOSE 8080

CMD ["./api_main"]

FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go build -o command_executor_main ./command-executor/cmd/app/main.go

EXPOSE 8082

CMD ["./command_executor_main"]

FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go build -o query_handler_main ./query-handler/cmd/app/main.go

EXPOSE 8081

CMD ["./query_handler_main"]
