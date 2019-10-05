# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:1.13.1 as builder

COPY . .
RUN go build -o /bin/app

FROM alpine:3.10

WORKDIR /bin/

COPY --from=builder /bin/app .

ENTRYPOINT ["/bin/app"]
