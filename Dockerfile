FROM golang:1.22-alpine

ENV GOPATH=/

COPY ./ ./

RUN go get -d ./...

CMD ["go", "run", "./cmd/"]
