# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:1.13.1 as builder

RUN mkdir -p "$GOPATH/src/github.com/1995parham/loser-scraper"
WORKDIR $GOPATH/src/github.com/1995parham/loser-scraper
ENV GO111MODULE=on

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app

FROM alpine:3.10

WORKDIR /bin/

COPY --from=builder /bin/app .

ENTRYPOINT ["/bin/app"]
