FROM golang:1.16-alpine

WORKDIR /album

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /out/go-sample-app

EXPOSE 8080

CMD ["./out/go-sample-app"]



