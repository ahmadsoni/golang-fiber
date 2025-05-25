FROM golang:1.24

WORKDIR /app

# Install psql client
RUN apt-get update && apt-get install -y postgresql-client

# Install air dengan path yang baru
RUN go install github.com/air-verse/air@latest

# Install migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
  mv migrate /usr/local/bin/migrate


COPY . .

RUN go mod tidy

CMD ["air"]
