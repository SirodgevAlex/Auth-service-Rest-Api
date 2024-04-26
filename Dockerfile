FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]


# FROM golang:1.17-alpine AS build

# WORKDIR /app

# COPY . .

# RUN go build -o main .

# FROM alpine:latest

# RUN apk update

# COPY --from=build /app/main /app/main

# WORKDIR /app

# CMD ["./main"]


# FROM golang:latest

# WORKDIR /app

# COPY ./go-notes-service .

# RUN go mod download

# RUN go build -o main ./cmd

# EXPOSE 8080

# CMD ["./main"]