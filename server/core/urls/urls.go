package urls

import (
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskHandlers "server/app_kinopoisk"
	userHandlers "server/app_user"
	filmsHandlers "server/app_films"

	"server/core/services"
)


func InitRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	// юзеры
	userGroup := apiV1.Group("/user")
	userGroup.Post("/sign-up", userHandlers.SignUp)
	userGroup.Post("/login", userHandlers.Login)
	userGroup.Use(services.AccessTokenMiddleware)
	userGroup.Use(services.BlacklistedTokenMiddleware)
	userGroup.Post("/logout", userHandlers.Logout)

	// фильмы от Kinopoisk API
	kinopoiskGroup := apiV1.Group("/kinopoisk/films")
	kinopoiskGroup.Use(services.AccessTokenMiddleware)
	kinopoiskGroup.Use(services.BlacklistedTokenMiddleware)
	kinopoiskGroup.Get("/search", kinopoiskHandlers.SearchFilms)
	kinopoiskGroup.Get("/:kinopoiskID", kinopoiskHandlers.GetFilmInfo)
	
	// сохранённые в БД фильмы
	filmsGroup := apiV1.Group("/films")
	filmsGroup.Use(services.AccessTokenMiddleware)
	filmsGroup.Use(services.BlacklistedTokenMiddleware)
	filmsGroup.Post("/:movieID/star", filmsHandlers.Star)
	filmsGroup.Post("/:movieID/unstar", filmsHandlers.Unstar)
	filmsGroup.Get("/stared", filmsHandlers.Stared)
	filmsGroup.Post("/:movieID/want", filmsHandlers.SetWant)
	filmsGroup.Post("/:movieID/watched", filmsHandlers.SetWatched)
	filmsGroup.Post("/:movieID/clear", filmsHandlers.Clear)
	filmsGroup.Get("/want", filmsHandlers.Want)
	filmsGroup.Get("/watched", filmsHandlers.Watched)
}
