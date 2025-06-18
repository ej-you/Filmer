# Filmer movies syncronizer

Movie-sync is consumer of RabbitMQ broker. It subscribes to update movie messages.
This messages are published by main REST API (server) and contains movie ID.
Movie-sync gets message, parses movie ID and send request to main REST API (server) to
update movie info.

> It's supposed to run `movie-sync` in a closed network.
> So there is no auth.

## Needed `env` variables

```dotenv
REST_API_HOST="https://example.com"
```

> Also used env vars for [RabbitMQ](../rabbitmq/README.md)
