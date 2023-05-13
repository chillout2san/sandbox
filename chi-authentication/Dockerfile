FROM golang:1.19-alpine

WORKDIR /app

COPY ./api .

RUN go mod tidy \
    && go build ./cmd/main.go

CMD ["./main"]