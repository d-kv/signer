FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go build -o command_executor_main ./command-executor/cmd/app/main.go

EXPOSE 8082

CMD ["./command_executor_main"]
