FROM golang:1.17 as builder

WORKDIR /go/src/github.com/appvia/certificate-monitor

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go install -v github.com/appvia/certificate-monitor

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -S app -G app

USER 1000

COPY --from=builder /go/bin/certificate-monitor /app

CMD ["/app"]
