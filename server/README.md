# Filmer server


### Needed `env` variables:

```
SERVER_PORT=3000
JWT_SECRET="samplelrhksgvi8n54kJWTgl58ehvyooSECREThielghi"

SERVER_CORS_ALLOWED_ORIGINS="*"
SERVER_CORS_ALLOWED_METHODS="GET"

# Key from Kinopoisk API Unofficial
KINOPOISK_API_UNOFFICIAL_KEY="jgw48oh5-g4n79-gn57wb9-fh643o78gwgj5"
# Key from Kinopoisk API
KINOPOISK_API_KEY="BR6BU5N6-5Y6ER-NJHES4-SW4MSIM6UBBR"

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
