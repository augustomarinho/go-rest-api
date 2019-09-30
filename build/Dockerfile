# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.13.1-buster as builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN mkdir -p /go/src/internal/app
WORKDIR /go/src/internal

COPY ./Gopkg.toml ./Gopkg.lock ./
COPY ./app/ ./app/
RUN dep ensure -vendor-only

WORKDIR /go/src/internal/app
RUN ls
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/internal/app .

CMD ["/internal/app/main"]