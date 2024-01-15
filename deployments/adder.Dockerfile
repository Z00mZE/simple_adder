FROM golang:alpine as Builder
WORKDIR /build
COPY . .
RUN apk add git ca-certificates
RUN go mod download  \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w" -o service ./cmd/calc/app.go

FROM alpine
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
COPY --from=Builder /build/service /service
ENTRYPOINT ["/service"]
