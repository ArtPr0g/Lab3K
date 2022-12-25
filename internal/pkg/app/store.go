package app

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// GetFilms godoc
// @Summary Get all stores
// @Description Get a list of all stores
// @Tags Info
// @Produce json
// @Success 200 {object} ds.StoreDocs
// @Failure 500 {object} swagger.Error
// @Router /film [get]
func (a *Application) GetFilms(gCtx *gin.Context) {
	resp, err := a.repo.GetFilms()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Can't get a list of promo codes",
				Error:       models.Err500,
				Type:        models.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetFilm godoc
// @Summary Get film
// @Description Get store using its uuid
// @Tags Info
// @Produce json
// @Param UUID path string true "UUID магазина" format(uuid)
// @Success 200 {object} ds.StoreDocs
// @Failure 500 {object} swagger.Error
// @Router /film/{UUID} [get]
func (a *Application) GetFilm(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	resp, err := a.repo.GetFilm(UUID)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Can't get a list of promo codes",
				Error:       models.Err500,
				Type:        models.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// ChangeFilm		godoc
// @Summary      	Changefilm
// @Description  	Change the film using its uuid
// @Tags         	Change
// @Produce      	json
// @Param 			UUID path string true "UUID фильма" format(uuid)
// @Param 			Price body ds.PriceFilm true "Новая цена"
// @Success      	200 {object} models.ModelPriceChanged
// @Failure 		400 {object} models.ModelError
// @Failure 		404 {object} models.ModelError
// @Failure 	 	500 {object} models.ModelError
// @Router       	/film/{UUID} [put]
func (a *Application) ChangeFilm(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       models.Err400,
				Type:        models.TypeClientReq,
			})
		return
	}

	film := ds.Film{}
	err = gCtx.BindJSON(&film)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price is negative or not int",
				Error:       models.Err400,
				Type:        models.TypeClientReq,
			})
		return
	}

	code, err := a.repo.ChangeFilm(UUID, film)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       models.Err404,
					Type:        models.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Change failed",
					Error:       models.Err500,
					Type:        models.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelPriceChanged{
			Success: true,
		})
}

// DeleteFilm   godoc
// @Summary      Delete a film
// @Description  Delete a film via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID фильмы"
// @Success      200  {object}  models.ModelFilmDeleted
// @Failure 	 500 {object} models.ModelError
// @Router       /film [delete]
func (a *Application) DeleteFilm(gCtx *gin.Context) {
	uuid, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       models.Err400,
				Type:        models.TypeClientReq,
			})
		return
	}

	code, err := a.repo.DeleteFilm(uuid)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       models.Err404,
					Type:        models.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Delete failed",
					Error:       models.Err500,
					Type:        models.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelFilmDeleted{
			Success: true,
		})
}

// CreateFilm godoc
// @Summary      Add a new film
// @Description  Adding a new film to database
// @Tags         Add
// @Produce      json
// @Param Name query string true "Название фильма"
// @Param Release query uint64 true "Дата выхода фильма"
// @Param Grade  query float64 true "Оценка фильма"
// @Param Genre query string true "Жанр фильма"
// @Param Price query int true "Стоимоть фильма"
// @Param WhatchTime query uint64 true "Длительность фильма(мин.)"
// @Param Summary  query string true "Описание"
// @Param Image query string true "Фото"
// @Param Video query string true "Видео"
// @Success      201  {object}  models.ModelFilmCreated
// @Failure 500 {object} models.ModelError
// @Router       /film [Post]
func (a *Application) CreateFilm(gCtx *gin.Context) {
	film := ds.Film{}

	if err := gCtx.BindJSON(&film); err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	if film.Price < 0 {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price cannot be non -negative",
				Error:       "Price error",
				Type:        "client",
			})
		return
	}
	if film.Grade < 0 || film.Grade > 10 {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price cannot be non -negative or more than 10",
				Error:       "Grade error",
				Type:        "client",
			})
		return
	}
	if film.WhatchTime < 0 {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The time cannot be non -negative",
				Error:       "WhatchTime error",
				Type:        "client",
			})
		return
	}
	film.UUID = uuid.New()
	err := a.repo.CreateFilm(film)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelFilmCreated{
			Success: true,
		})
}

// GetVideo			godoc
// @Summary     	Get a promo
// @Description 	Get a promo in store using its uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID path string true "UUID магазина" format(uuid)
// @Param 			Quantity path string true "Кол-во"
// @Success      	200 {object} swagger.StorePromo
// @Failure 		400 {object} swagger.Error
// @Failure 		404 {object} swagger.Error
// @Failure 	 	500 {object} swagger.Error
// @Router       	/film/{UUID}/{Quantity} [get]
func (a *Application) GetVideo(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	FilmUUID, err := uuid.Parse(gCtx.Param("uuid"))
	quantity, _ := strconv.Atoi(gCtx.Param("quantity"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			models.ModelError{
				Description: "Invalid UUID format",
				Error:       models.Err400,
				Type:        models.TypeClientReq,
			})
		return
	}

	code, Video, err := a.repo.GetVideo(uint64(quantity), FilmUUID, userUUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       models.Err404,
					Type:        models.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Delete failed",
					Error:       models.Err500,
					Type:        models.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(http.StatusOK, Video)
}
