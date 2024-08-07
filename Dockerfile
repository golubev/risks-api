FROM golang:1.22-alpine

ENV GOPATH=/

COPY ./ ./

RUN go get -d ./...

RUN go build -o risks-api ./cmd/main.go

CMD ["./risks-api"]
