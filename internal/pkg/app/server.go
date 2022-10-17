package app

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/app/ds"
	"awesomeProject/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.New()

	r.GET("/films", a.GetList)
	r.GET("/films/price/:uuid", a.GetFilmPrice)

	r.POST("/films/create", a.AddFilm)

	r.PUT("/films/price/change/:uuid", a.ChangePrice)

	r.DELETE("/films/delete/:uuid", a.DeleteFilm)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

type inter struct {
	Status string `json:"status"`
}

// GetList godoc
// @Summary      Get all records
// @Description  Get a list of all films
// @Tags         Info
// @Produce      json
// @Success      200  {object}  ds.Kino
// @Failure 500 {object} models.ModelError
// @Router       /films [get]
func (a *Application) GetList(gCtx *gin.Context) {
	resp, err := a.repo.GetFilmList()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

// GetFilmPrice  godoc
// @Summary      Get price for a film
// @Description  Get a price via uuid of a film
// @Tags         Info
// @Produce      json
// @Param UUID query string true "UUID фильма"
// @Success      200  {object}  models.ModelFilmPrice
// @Failure 	 500 {object} models.ModelError
// @Router       /films/price [get]
func (a *Application) GetFilmPrice(gCtx *gin.Context) {
	uuid := gCtx.Param("uuid")
	log.Println(uuid)
	respName, respPrice, err := a.repo.GetFilmPrice(uuid)
	if err != nil {
		if respName == "no film found with this uuid" {
			gCtx.JSON(
				http.StatusBadRequest,
				&models.ModelError{
					Description: "No film found with this uuid",
					Error:       "uuid error",
					Type:        "client",
				})
			return
		}
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a price",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelFilmPrice{
			Name:  respName,
			Price: strconv.Itoa(respPrice),
		})
}

// ChangePrice   godoc
// @Summary      Change film price
// @Description  Change a price for a film via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID фильма"
// @Param Price query int true "Новая цена"
// @Success      200  {object}  models.ModelPriceChanged
// @Failure 	 500 {object} models.ModelError
// @Router       /films/price/change [put]
func (a *Application) ChangePrice(gCtx *gin.Context) {
	newPrice, _ := strconv.ParseUint(gCtx.Query("Price"), 10, 64)
	if newPrice < 0 {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price cannot be non -negative",
				Error:       "Price error",
				Type:        "client",
			})
		return
	}
	uuidR := gCtx.Param("uuid")
	inputUuid, _ := uuid.Parse(uuidR)
	err, messageError := a.repo.ChangePrice(inputUuid, newPrice)
	if err != nil {
		if messageError == "record not found" {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "record failed",
					Error:       "db error",
					Type:        "client",
				})
			return
		}
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Update failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
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
// @Router       /films/delete [delete]
func (a *Application) DeleteFilm(gCtx *gin.Context) {
	uuid := gCtx.Param("uuid")
	messageError, err := a.repo.DeleteFilm(uuid)
	if err != nil {
		if messageError == "no film found with this uuid" {
			gCtx.JSON(
				http.StatusBadRequest,
				&models.ModelError{
					Description: "No film found with this uuid",
					Error:       "uuid error",
					Type:        "client",
				})
			return
		}
		gCtx.JSON(
			http.StatusNotFound,
			&models.ModelError{
				Description: "delete failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelFilmDeleted{
			Success: true,
		})

}

// AddFilm godoc
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
// @Success      201  {object}  models.ModelFilmCreated
// @Failure 500 {object} models.ModelError
// @Router       /films/create [Post]
func (a *Application) AddFilm(gCtx *gin.Context) {
	film := ds.Kino{}

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
	film.UUID = uuid.New()
	err := a.repo.AddFilm(film)
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
