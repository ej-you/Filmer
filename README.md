# Filmer

## About [server](./server/README.md)

## About [client](./client/README.md)

## About [Nginx](./nginx/README.md)

## About [admin-panel](./admin/README.md)

## About [Rabbit MQ](./rabbitmq/README.md)

## About [movie-sync](./movie_sync/README.md)

### TODO

- [x] Change interface language to Russian (`ru_locale` branch)
- [x] Fix frontend bugs
- [x] Add `back` button to movie page
- [x] Add limit for movie directors amount (max = 8 directors)
- [x] Save cache data on redis shutdown (and load cache data on redis startup)
- [x] Add log rotation for server
- [x] Add to user category pages filter by substring
- [x] Add director/actor page with his movies
- [x] Add `step-down` flag for one step migrations rollback to migrator
- [x] Add an admin panel to monitor user activity
- [ ] Fix frontend bug: hyphenation of long words on the phone
- [ ] Add related movies to movie page
- [ ] Add popular movies (or recomendations or random "daily" movie) to the main page
- [ ] Add the ability to recomend the movie to another user
- [ ] Add the ability to browse recomended movies from other users
- [ ] Add OpenID auth
