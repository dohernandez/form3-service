FROM golang:1.11 AS builder

ARG VERSION=dev
ARG USER=dohernandez

WORKDIR /go/src/github.com/dohernandez/form3-service

# Install migrate
RUN  curl -sL https://github.com/golang-migrate/migrate/releases/download/v4.2.4/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /bin/migrate

COPY . .

RUN make build

FROM ubuntu:bionic

RUN groupadd -r dohernandez && useradd --no-log-init -r -g dohernandez dohernandez
USER dohernandez

COPY --from=builder --chown=dohernandez:dohernandez /go/src/github.com/dohernandez/form3-service/bin/form3-service /bin/form3-service
COPY --from=builder --chown=hellofresh:hellofresh /bin/migrate /bin/migrate

COPY resources/migrations /resources/migrations

EXPOSE 8000
ENTRYPOINT ["form3-service"]
