# Filmer admin-panel

It's supposed to run `admin-panel` in a closed network.
So there is no auth.

## Needed `env` variables

```dotenv
ADMIN_PANEL_PORT=3002
KEEP_ALIVE_TIMEOUT="5s"
GIN_MODE="release"

REST_API_HOST="https://example.com"
```
