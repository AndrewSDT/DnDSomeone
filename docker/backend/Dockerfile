FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

FROM alpine

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]