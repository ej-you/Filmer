package http

import (
	"gorm.io/gorm"
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/middlewares"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/validator"
	"Filmer/server/config"

	"Filmer/server/internal/auth"
	"Filmer/server/internal/auth/usecase"
	"Filmer/server/internal/auth/repository"
)


// тип интерфейса auth.Router
type authRouter struct {
	validator 		validator.Validator
	mwManager 		middlewares.MiddlewareManager
	authRepo 		auth.Repository
	authCacheRepo	auth.CacheRepository
	authUsecase		auth.Usecase
}

// конструктор для типа интерфейса auth.Router
func NewAuthRouter(
	cfg *config.Config,
	dbClient *gorm.DB,
	cacheClient cache.Cache,
	validator validator.Validator,
	mwManager middlewares.MiddlewareManager,
) auth.Router {
	authRepo := repository.NewRepository(dbClient)
	authCacheRepo := repository.NewCacheRepository(cfg, cacheClient)
	authUsecase := usecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &authRouter{
		validator: validator,
		mwManager: mwManager,
		authRepo: authRepo,
		authCacheRepo: authCacheRepo,
		authUsecase: authUsecase,
	}
}

func (this authRouter) SetRoutes(router fiber.Router) {
	router.Post("/sign-up", SignUp(this.authUsecase, this.validator))
	router.Post("/login", Login(this.authUsecase, this.validator))

	restricted := router.Use(this.mwManager.JWTAuth())
	restricted.Post("/logout", Logout(this.authUsecase))
}
