# Filmer nginx

* Set up to use in prod mode

To use it in dev mode, uncomment the next code in [http_site.conf](./http_site.conf):

```conf
# ...

# REST API (dev)
upstream backend {
    server server:3000;
}

# ...
# server
# ...

    # forward to REST API (dev)
    location /api/ {
        proxy_pass http://backend;
    }

# ...
```
