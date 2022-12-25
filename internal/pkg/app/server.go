package app

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/app/role"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *Application) StartServer() {
	log.Println("server start up")

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Запрос для свагера
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запросы для магазина:
	r.GET("/film", a.GetFilms)

	r.GET("/film/:uuid", a.GetFilm)

	r.GET("/film/:uuid/:quantity", a.GetVideo)

	// Запросы для корзины:
	r.GET("/cart/:film", a.GetCart1)

	r.GET("/cart/increase/:film", a.IncreaseQuantity)

	r.DELETE("/cart/delete/:film", a.DeleteCart)

	// Запросы для авторизации
	r.POST("/login", a.Login)

	r.POST("/sign_up", a.Register)

	r.GET("/logout", a.Logout)

	r.GET("/role", a.Role)

	// Запросы для всех авторизированных пользователей
	r.Use(a.WithAuthCheck(role.Buyer, role.Manager, role.Admin)).GET("/cart", a.GetCart)

	// Запросы для менеджеров
	r.Use(a.WithAuthCheck(role.Manager)).POST("/film", a.CreateFilm)

	r.Use(a.WithAuthCheck(role.Manager)).DELETE("/film/:uuid", a.DeleteFilm)

	r.Use(a.WithAuthCheck(role.Manager)).PUT("/film/:uuid", a.ChangeFilm)

	r.Use(a.WithAuthCheck(role.Manager)).GET("/orders", a.GetOrders)

	r.Use(a.WithAuthCheck(role.Manager)).PUT("/orders/:uuid", a.ChangeStatus)

	r.Use(a.WithAuthCheck(role.Manager)).GET("/user/:uuid", a.GetUser)

	_ = r.Run()

	log.Println("server down")
}
