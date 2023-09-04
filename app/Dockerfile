FROM golang:1.21

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/99designs/gqlgen@latest
