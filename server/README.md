# Filmer server


### Needed `env` variables:

```
SERVER_PORT=3000
JWT_SECRET="samplelrhksgvi8n54kJWTgl58ehvyooSECREThielghi"

SERVER_CORS_ALLOWED_ORIGINS="*"
SERVER_CORS_ALLOWED_METHODS="GET"

# Key from Kinopoisk API Unofficial
KINOPOISK_API_KEY="jgw48oh5-g4n79-gn57wb9-fh643o78gwgj5"
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
