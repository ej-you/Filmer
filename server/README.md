# Filmer server


### Server part use `clean architecture` described in [README_ARCH](./docs/README_ARCH.md)
### App description can be found [there](./docs/README_APP.md)

### Needed `env` variables:

```
SERVER_PORT=3000
JWT_SECRET="samplelrhksgvi8n54kJWTgl58ehvyooSECREThielghi"

SERVER_CORS_ALLOWED_ORIGINS="*"
SERVER_CORS_ALLOWED_METHODS="GET,POST"

# JWT token expired duration
TOKEN_EXPIRED="1h"
# Kinopoisk movie data expired duration
KINOPOISK_API_DATA_EXPIRED="360h"

# Key from Kinopoisk API Unofficial
KINOPOISK_API_UNOFFICIAL_KEY="sample8h5-g4n79-gn57wb9-fh643o78gwgj5"
# Key from Kinopoisk API
KINOPOISK_API_KEY="SAMPLEU5N6-5Y6ER-NJHES4-SW4MSIM6UBBR"

REDIS_HOST=172.21.0.3
REDIS_PORT=6379

# CockroachDB used in insecure mode
DB_USER="root"
DB_HOST=172.21.0.2
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
2. Migrate DB (x2 one-step up)
```
make migrate-up
make migrate-up
```
3. Up full app
```
docker compose up -d
```

<hr>

### Used tools:

1. [CockroachDB](https://www.cockroachlabs.com/) as main DB
2. DB migrations with [migrate module](https://github.com/golang-migrate/migrate)
3. [Redis](https://github.com/redis/go-redis) as cache
4. [Fiber](https://docs.gofiber.io/) for RESTful API server
5. Structs validation with [go-playground validator](https://github.com/go-playground/validator)
6. JSON (de)serializer with [easyjson](https://github.com/mailru/easyjson)
7. JWT as access token for user session
8. Swagger docs with [swaggo v2](https://github.com/swaggo/swag)
9. Golang linter - [golangci-lint](https://golangci-lint.run/)

### Third party:

1. Kinopoisk API Key from [Kinopoisk API Unofficial](https://kinopoiskapiunofficial.tech/)
2. Kinopoisk API Key from [Kinopoisk API](https://kinopoisk.dev//)
