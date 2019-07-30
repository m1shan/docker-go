FROM golang:alpine as builder

RUN apk add --update gcc g++ git curl

WORKDIR /go/src/app

COPY . .

RUN go get ./...
RUN GOOS=linux go build -o ./bin/server ./main.go

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin

COPY --from=builder /go/src/app/bin /go/bin

EXPOSE 8080
ENTRYPOINT /go/bin/server --port 8080
