FROM golang:1.11 AS builder

ARG VERSION=dev
ARG USER=dohernandez

WORKDIR /go/src/github.com/dohernandez/form3-service

COPY . .

RUN make build

FROM ubuntu:bionic

RUN groupadd -r dohernandez && useradd --no-log-init -r -g dohernandez dohernandez
USER dohernandez

COPY --from=builder --chown=dohernandez:dohernandez /go/src/github.com/dohernandez/form3-service/bin/form3-service /bin/form3-service

EXPOSE 8000
ENTRYPOINT ["form3-service"]
