# Filmer Rabbit MQ

Rabbit MQ used for background updates of movie info

## Needed `env` variables (in file `/rabbitmq/.env`)

```dotenv
# set user name and password
RABBITMQ_DEFAULT_USER=rabbit
RABBITMQ_DEFAULT_PASS=rabbit

# connection settings for producers/consumers (AMQP)
RABBITMQ_HOST=127.0.0.1
RABBITMQ_PORT=5672
```
