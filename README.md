# enechain Technical Assignment

## Setup

```sh
# Create databases
docker compose up -d postgres
docker compose exec psql --user postgres --command 'CREATE DATABASE app'
docker compose exec psql --user postgres --command 'CREATE DATABASE app_test'

# Create tables
docker compose run --rm web make migrate
```

## Development

```sh
# Run public web server
docker compose up web

# Run admin web server
docker compose up admin
```
