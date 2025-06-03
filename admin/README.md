# Filmer admin-panel

It's supposed to run `admin-panel` in a closed network.
So there is only one admin. Username and password for him should be specified in ENV-variables.
By default it's "admin" - "admin".

## Needed `env` variables

```dotenv
ADMIN_PANEL_PORT=3001

REST_API_HOST="https://example.com"
TOKEN_EXPIRED="1m"

ADMIN_USERNAME="admin"
ADMIN_PASSWORD="admin"
```
