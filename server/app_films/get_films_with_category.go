package app_films

import (
	"fmt"
	"math"

	"github.com/google/uuid"
	"gorm.io/gorm"
	fiber "github.com/gofiber/fiber/v2"

	coreValidator "server/core/validator"
	"server/core/services"
	"server/db/schemas"
	"server/db"
)


// категории выборки фильмов юзера
const (
	selectStaredCategory = "stared = true" // избранное
	selectWantCategory = "status = 1" // хочу посмотреть
	selectWatchedCategory = "status = 2" // посмотрел
	
	paginationLimit = 10 // сколько записей будет возвращаться на одной странице
)



//easyjson:json
type Filter struct {
	RatingFrom	*float64 `json:"ratingFrom,omitempty" query:"ratingFrom" validate:"omitempty,min=0"`
	YearFrom	int `json:"yearFrom,omitempty" query:"yearFrom" validate:"omitempty,min=1500,max=3000"`
	YearTo		int `json:"yearTo,omitempty" query:"yearTo" validate:"omitempty,min=1500,max=3000"`
	Type		string `json:"type,omitempty" query:"type" validate:"omitempty,oneof=фильм сериал видео мини-сериал"`
	Genres		[]string `json:"genres,omitempty" query:"genres" validate:"omitempty"`
}
//easyjson:json
type Sort struct {
	SortField 	string `json:"sortField,omitempty" query:"sortField" validate:"omitempty,oneof=title rating year updated_at"`
	SortOrder 	string `json:"sortOrder,omitempty" query:"sortOrder" validate:"omitempty,oneof=asc desc"`
}
//easyjson:json
type Pagination struct {
	Page int 	`json:"page" query:"page" validate:"omitempty,min=1"`
	Pages int 	`json:"pages"`
	Total int64 `json:"total"`
	Limit int 	`json:"limit"`
}

//easyjson:json
type CategoryFilmsIn struct {
	// поля для фильтрации
	Filter
	// сортировка по полю
	Sort
	// пагинация
	Pagination
	// uuid юзера
	UserID uuid.UUID `json:"-"`
}

//easyjson:json
type CategoryFilmsOut struct {
	// поля для фильтрации
	Filter		Filter `json:"filter"`
	// сортировка по полю
	Sort		Sort `json:"sort"`
	// пагинация
	Pagination	Pagination `json:"pagination"`
	// список найденных фильмов
	Films		schemas.UserMovies `json:"films"`
}


// получение фильмов юзера определённой категории из БД
func doSelect(dbConn *gorm.DB, filmsCategory string, dataIn *CategoryFilmsIn, userMovies *schemas.UserMovies) *gorm.DB {
	// базовые параметры
	selectQuery := dbConn.
		Table("user_movies").
		Distinct("user_movies.*, movies.title, movies.rating, movies.year").
		InnerJoins("INNER JOIN movies ON user_movies.movie_id = movies.id").
		Where(filmsCategory).
		Where("user_id = ?", dataIn.UserID)
	// фильтры
	selectQuery = addSort(selectQuery, &dataIn.Sort)
	// сортировка
	selectQuery = addFilter(selectQuery, &dataIn.Filter)
	// пагинация
	selectQuery = addPagination(selectQuery, &dataIn.Pagination)
	// добавляем подгрузку данных из зависимых таблиц
	selectQuery = selectQuery.
		Preload("Movie", func(dbOmit *gorm.DB) *gorm.DB {
		    // исключаем тяжёлые поля из подгрузки данных
		    return dbOmit.Omit("WebURL", "FilmLength", "Description", "Personal")
		}).
		Preload("Movie.Genres")

	return selectQuery.Find(&userMovies)
}
// добавление клиентских фильтров к выборке
func addFilter(dbConn *gorm.DB, filter *Filter) *gorm.DB {
	// с какого рейтинга
	if filter.RatingFrom != nil {
		dbConn = dbConn.Where("rating >= ?", filter.RatingFrom)
	}
	// с какого года
	if filter.YearFrom != 0 {
		dbConn = dbConn.Where("year >= ?", filter.YearFrom)
	}
	// до какого года
	if filter.YearTo != 0 {
		dbConn = dbConn.Where("year <= ?", filter.YearTo)
	}
	// тип фильма
	if filter.Type != "" {
		dbConn = dbConn.Where("type = ?", filter.Type)
	}
	// имеет один из жанров 
	if len(filter.Genres) > 0 {
		dbConn = dbConn.
			InnerJoins(`INNER JOIN genres ON genres.movie_id = movies.id`).
			Where(`genres.genre IN ?`, filter.Genres)
	}
	return dbConn
}
// добавление клиентской сортировки к выборке
func addSort(dbConn *gorm.DB, sort *Sort) *gorm.DB {
	// если поле сортировки не задано, то сортируем по дате обновления
	if sort.SortField == "" {
		sort.SortField = "updated_at"
	}

	// если при неуказанном поле сортировки (или updated_at) не задан порядок, то ставим обратный
	if sort.SortField == "updated_at" && sort.SortOrder == "" {
		sort.SortOrder = "desc"
	} else {
		// если при указанном поле сортировки не задан порядок, то ставим прямой
		if sort.SortOrder == "" {
			sort.SortOrder = "asc"
		}
	}
	return dbConn.Order(fmt.Sprintf("%s %s", sort.SortField, sort.SortOrder))
}
// добавление данных пагинации к выборке (добавлять после всех фильтров)
func addPagination(dbConn *gorm.DB, pagination *Pagination) *gorm.DB {
	// если страница задана, то выставляем первую
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	// установка лимита записей на страницу
	pagination.Limit = paginationLimit
	
	// получение кол-ва всех записей, подходящих под фильтры
	dbConn.Count(&pagination.Total)
	// высчитывание кол-ва страниц
	pagination.Pages = int(math.Ceil(float64(pagination.Total) / float64(pagination.Limit)))

	return dbConn.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)
}


// получение фильмов определённой категории
func getFilmsWithCategory(ctx *fiber.Ctx, filmsCategory string) error {
	var err error
	var dataIn = CategoryFilmsIn{UserID: services.ParseUserIDFromContext(ctx)}
	var dataOut CategoryFilmsOut

	// парсинг query-параметров
	if err = ctx.QueryParser(&dataIn); err != nil {
		return fmt.Errorf("get films where %s: %w", filmsCategory, err)
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return fmt.Errorf("get films where %s: %w", filmsCategory, err)
	}

	// получение избранных фильмов юзера из БД
	dbConn := db.GetConn()
	var userMovies schemas.UserMovies
	selectResult := doSelect(dbConn, filmsCategory, &dataIn, &userMovies)
	if err = selectResult.Error; err != nil {
		return fmt.Errorf("get films where %s: %w", filmsCategory, fiber.NewError(500, "failed to get user movies: " + err.Error()))
	}

	dataOut.Filter = dataIn.Filter
	dataOut.Sort = dataIn.Sort
	dataOut.Pagination = dataIn.Pagination
	dataOut.Films = userMovies
	return ctx.Status(200).JSON(dataOut)
}


//	@summary		Получение избранных фильмов юзера
//	@description	Получение избранных фильмов юзера с пагинацией и настраиваемой сортировкой и фильтрацией
//	@router			/films/stared [get]
//	@id				films-get-stared
//	@tags			films
//	@security		JWT
//	@param			page		query		int			false	"страница поиска (Например: 1)"
//	@param			sortField	query		string		false	"поле для сортировки (Например: year | По умолчанию: updated_at | Допустимые значения: title, rating, year, updated_at)"
//	@param			sortOrder	query		string		false	"направление сортировки (Например: desc | По умолчанию: asc [для updated_at: desc] | Допустимые значения: asc, desc)"
//	@param			ratingFrom	query		float64		false	"минимальный рейтинг (Например: 7.5 | Допустимые значения: 0 и больше)"
//	@param			yearFrom	query		int			false	"минимальный год (Например: 1991 | Допустимые значения: 1500..3000)"
//	@param			yearTo		query		int			false	"максимальный год (Например: 2014 | Допустимые значения: 1500..3000)"
//	@param			type		query		string		false	"тип фильма (Например: сериал | Допустимые значения: фильм, сериал, видео, мини-сериал)"
//	@param			genres		query		[]string	false	"жанры фильмов (Например: боевик)"
//	@success		200			{object}	schemas.UserMovie
//	@failure		401			"Пустой или неправильный токен"
//	@failure		403			"Истекший или невалидный токен"
func Stared(ctx *fiber.Ctx) error {
	return getFilmsWithCategory(ctx, selectStaredCategory)
}

//	@summary		Получение фильмов юзера из списка "хочу посмотреть"
//	@description	Получение фильмов юзера из списка "хочу посмотреть" с пагинацией и настраиваемой сортировкой и фильтрацией
//	@router			/films/want [get]
//	@id				films-get-want
//	@tags			films
//	@security		JWT
//	@param			page		query		int			false	"страница поиска (Например: 1)"
//	@param			sortField	query		string		false	"поле для сортировки (Например: year | По умолчанию: updated_at | Допустимые значения: title, rating, year, updated_at)"
//	@param			sortOrder	query		string		false	"направление сортировки (Например: desc | По умолчанию: asc [для updated_at: desc] | Допустимые значения: asc, desc)"
//	@param			ratingFrom	query		float64		false	"минимальный рейтинг (Например: 7.5 | Допустимые значения: 0 и больше)"
//	@param			yearFrom	query		int			false	"минимальный год (Например: 1991 | Допустимые значения: 1500..3000)"
//	@param			yearTo		query		int			false	"максимальный год (Например: 2014 | Допустимые значения: 1500..3000)"
//	@param			type		query		string		false	"тип фильма (Например: сериал | Допустимые значения: фильм, сериал, видео, мини-сериал)"
//	@param			genres		query		[]string	false	"жанры фильмов (Например: боевик)"
//	@success		200			{object}	schemas.UserMovie
//	@failure		401			"Пустой или неправильный токен"
//	@failure		403			"Истекший или невалидный токен"
func Want(ctx *fiber.Ctx) error {
	return getFilmsWithCategory(ctx, selectWantCategory)
}

//	@summary		Получение фильмов юзера из списка "посмотрел"
//	@description	Получение фильмов юзера из списка "посмотрел" с пагинацией и настраиваемой сортировкой и фильтрацией
//	@router			/films/watched [get]
//	@id				films-get-watched
//	@tags			films
//	@security		JWT
//	@param			page		query		int			false	"страница поиска (Например: 1)"
//	@param			sortField	query		string		false	"поле для сортировки (Например: year | По умолчанию: updated_at | Допустимые значения: title, rating, year, updated_at)"
//	@param			sortOrder	query		string		false	"направление сортировки (Например: desc | По умолчанию: asc [для updated_at: desc] | Допустимые значения: asc, desc)"
//	@param			ratingFrom	query		float64		false	"минимальный рейтинг (Например: 7.5 | Допустимые значения: 0 и больше)"
//	@param			yearFrom	query		int			false	"минимальный год (Например: 1991 | Допустимые значения: 1500..3000)"
//	@param			yearTo		query		int			false	"максимальный год (Например: 2014 | Допустимые значения: 1500..3000)"
//	@param			type		query		string		false	"тип фильма (Например: сериал | Допустимые значения: фильм, сериал, видео, мини-сериал)"
//	@param			genres		query		[]string	false	"жанры фильмов (Например: боевик)"
//	@success		200			{object}	schemas.UserMovie
//	@failure		401			"Пустой или неправильный токен"
//	@failure		403			"Истекший или невалидный токен"
func Watched(ctx *fiber.Ctx) error {
	return getFilmsWithCategory(ctx, selectWatchedCategory)
}
