FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk update

COPY --from=build /app/main /app/main

WORKDIR /app

CMD ["/app/main"]
