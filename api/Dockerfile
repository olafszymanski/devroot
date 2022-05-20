FROM golang:1.18-alpine AS build

RUN mkdir /api

WORKDIR /api

COPY . .

RUN go mod download

WORKDIR /api/cmd

RUN go build -o main



FROM alpine:latest

EXPOSE 3000

COPY --from=build /api/cmd/main .

CMD /main