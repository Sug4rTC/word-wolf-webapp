FROM golang:1.23-alpine

WORKDIR /app/backend 

COPY ./word-wolf-backend/go.mod ./word-wolf-backend/go.sum ./
RUN go mod download

COPY ./word-wolf-backend ./

RUN mkdir -p /app/bin
RUN go build -o /app/bin/main ./cmd/server

EXPOSE 8080

CMD ["/app/bin/main"]
