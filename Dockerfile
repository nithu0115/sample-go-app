FROM golang:latest
WORKDIR /go/src/github.com/nithu0115/sample-go-app
COPY . .
RUN CGO_ENABLED=0 go build -v -o main

FROM scratch
COPY --from=0 /go/src/github.com/nithu0115/sample-go-app/main .
ENTRYPOINT ["/main"]
