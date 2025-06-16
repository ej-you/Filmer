package kinopoisk

type CacheRepo interface {
	SetAPILimit(apiName string) error
	IsAPILimitReached(apiName string) (bool, error)
}
