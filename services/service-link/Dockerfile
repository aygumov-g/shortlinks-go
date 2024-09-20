FROM golang:1.22-alpine as builder

WORKDIR /src

COPY ./src/go.mod .
RUN go mod download

COPY ./src .

ARG PORT
ENV PORT $SERVER_PORT

EXPOSE $SERVER_PORT

RUN go build -o /main cmd/app/main.go

FROM alpine:3

COPY --from=builder main /bin/main

ENTRYPOINT ["/bin/main"]
