package app

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

/*func (a *Application) AddOrder(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)
	order := ds.Order{}
	order.UserUUID = userUUID
	err := gCtx.BindJSON(&order)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid parameters",
				Error:       models.Err400,
				Type:        models.TypeClientReq,
			})
		return
	}
	err = a.repo.AddOrder(order)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Create failed",
				Error:       models.Err500,
				Type:        models.TypeInternalReq,
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelFilmCreated{
			Success: true,
		})

}*/

func (a *Application) GetOrders(gCtx *gin.Context) {
	stDate := gCtx.Query("start_date")
	endDate := gCtx.Query("end_date")
	status := gCtx.Query("status")
	resp, err := a.repo.GetOrders(stDate, endDate, status)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
				Error:       models.Err500,
				Type:        models.TypeInternalReq,
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

func (a *Application) ChangeStatus(gCtx *gin.Context) {
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
	order := ds.Order{}
	err = gCtx.BindJSON(&order)
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
	resp, err := a.repo.ChangeStatus(UUID, order.Status)
	if err != nil {
		if resp == 404 {
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
		&models.ModelCartCreated{
			Success: true,
		})

}
