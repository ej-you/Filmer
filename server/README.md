# Filmer server


### Needed `env` variables:

```
SERVER_PORT=3000
JWT_SECRET="samplelrhksgvi8n54kJWTgl58ehvyooSECREThielghi"

SERVER_CORS_ALLOWED_ORIGINS="*"
SERVER_CORS_ALLOWED_METHODS="GET"

# Key from Kinopoisk API Unofficial
KINOPOISK_API_UNOFFICIAL_KEY="e1167bc7-1c81-4bf6-86cf-59d9be5adbe9"
# Key from Kinopoisk API
KINOPOISK_API_KEY="V01XTHC-2XQMAHD-G2BY8XH-1R6XWQH"

# CockroachDB used in insecure mode
DB_USER="root"
DB_HOST=172.18.0.2
DB_PORT=26257
DB_NAME="filmer_db"
```

<hr>

### Deploy

1. Create database
```
docker compose up -d cockroach
docker exec -it cockroach ./cockroach sql --insecure --execute="CREATE DATABASE IF NOT EXISTS filmer_db;"
```
2. Up full app
```
docker compose up -d
```


### Used tools:

1. Kinopoisk API Key from [Kinopoisk API Unofficial](https://kinopoiskapiunofficial.tech/)
2. Kinopoisk API Key from [Kinopoisk API](https://kinopoisk.dev//)
