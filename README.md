# Filmer

### About [server](./server/README.md)
### About [client](./client/README.md)
### About [nginx](./nginx/README.md)

<hr>

### TODO

- [x] Change interface language to Russian (`ru_locale` branch)
- [x] Fix frontend bugs (at least):
	- [x] Opening a modal not from the very top of the screen
	- [x] Remove movies cards stretch
	- [x] Set `width` for input `#search-keyword` in search movies
- [ ] Add popular movies (or recomendations or random "daily" movie) to the main page
- [ ] Add OpenID auth
- [ ] Add the ability to recomend the movie to another user
- [ ] Add the ability to browse recomended movies from other users
- [ ] Add director/actor page with his movies
- [ ] Add `back` button to movie page
- [ ] Add related movies to movie page
- [x] Add limit for movie directors amount (max = 8 directors)
- [x] Save cache data on redis shutdown (and load cache data on redis startup)
- [x] Add log rotation for server
- [ ] Add `step-down` flag for one step migrations rollback to migrator
