package app

import (
	"awesomeProject/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// GetCart 			godoc
// @Summary      	Get a whole cart
// @Description  	Get a list of the entire basket
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} ds.Cart
// @Failure 		500 {object} models.ModelError
// @Router       	/cart [get]
func (a *Application) GetCart(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	resp, err := a.repo.GetCart(userUUID)
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

// GetCart1 		godoc
// @Summary      	Get film from the cart
// @Description  	Get one film from the shopping cart
// @Tags         	Info
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} ds.Cart
// @Failure 		400 {object} models.ModelError
// @Failure 		500 {object} models.ModelError
// @Router       	/cart/{Film} [get]
func (a *Application) GetCart1(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	film, err := uuid.Parse(gCtx.Param("film"))
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

	resp, _ := a.repo.GetCart1(film, userUUID)

	gCtx.JSON(http.StatusOK, resp)
}

// IncreaseQuantity godoc
// @Summary      	Increase by 1 in the cart
// @Description  	Increase by 1 the number of promo codes in the cart
// @Tags         	Info
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} swagger.CartIncrease
// @Failure 		400 {object} models.ModelError
// @Failure 		500 {object} models.ModelError
// @Router       	/cart/increase/{Film} [get]
func (a *Application) IncreaseQuantity(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	film, err := uuid.Parse(gCtx.Param("film"))
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

	quantity, err := a.repo.IncreaseQuantity(film, userUUID)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Change failed",
				Error:       models.Err500,
				Type:        models.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, quantity)
}

// DeleteCart		godoc
// @Summary     	Delete a store in the cart
// @Description 	Delete a store in the cart using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			Store path string true "Магазин"
// @Success      	200 {object} swagger.Delete
// @Failure 		400 {object} models.ModelError
// @Failure 		404 {object} models.ModelError
// @Failure 	 	500 {object} models.ModelError
// @Router       	/cart/delete/{Film} [delete]
func (a *Application) DeleteCart(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	film, err := uuid.Parse(gCtx.Param("film"))
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

	code, err := a.repo.DeleteCart(film, userUUID)
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
