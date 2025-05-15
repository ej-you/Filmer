# Filmer

## About [server](./server/README.md)

## About [client](./client/README.md)

## About [nginx](./nginx/README.md)

### TODO

- [x] Change interface language to Russian (`ru_locale` branch)
- [x] Fix frontend bugs
- [x] Add `back` button to movie page
- [x] Add limit for movie directors amount (max = 8 directors)
- [x] Save cache data on redis shutdown (and load cache data on redis startup)
- [x] Add log rotation for server
- [x] Add to user category pages filter by substring
- [x] Add director/actor page with his movies
- [ ] Add related movies to movie page
- [ ] Add popular movies (or recomendations or random "daily" movie) to the main page
- [ ] Add the ability to recomend the movie to another user
- [ ] Add the ability to browse recomended movies from other users
- [ ] Add `step-down` flag for one step migrations rollback to migrator
- [ ] Add OpenID auth
- [ ] Add an admin panel to monitor user activity
