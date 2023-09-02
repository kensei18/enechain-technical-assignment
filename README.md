# enechain Technical Assignment

## Setup

```sh
docker compose up -d postgres
docker compose exec psql --user postgres --command 'CREATE DATABASE app'
docker compose exec psql --user postgres --command 'CREATE DATABASE app_test'
```
