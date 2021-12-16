FROM golang:1.13 as builder

WORKDIR /go/src/github.com/appvia/certificate-monitor

RUN go get github.com/golang/dep/cmd/dep
COPY go.mod go.sum ./

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go install -v github.com/appvia/certificate-monitor

FROM alpine:3.11.3

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -S app -G app

USER 1000

COPY --from=builder /go/bin/certificate-monitor /app

CMD ["/app"]
