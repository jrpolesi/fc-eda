FROM golang:1.20

COPY go.mod go.sum /app/

WORKDIR /app/

ENV DOCKERIZE_VERSION v0.8.0

RUN apt-get update \
    && apt-get install -y wget \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apt-get autoremove -yqq --purge wget && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y librdkafka-dev

RUN go mod tidy

CMD ["go", "run", "./cmd/api/main.go"]
