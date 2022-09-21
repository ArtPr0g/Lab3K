package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Users struct {
	Id      int     `json:"Id"`
	Name    string  `json:"Name"`
	Release int     `json:"Release"`
	Grade   float64 `json:"Grade"`
}

func StartServer() {
	log.Println("Server start up")

	var users = []Users{
		{Id: 1, Name: "Крепкий орешек", Release: 1988, Grade: 8.0},
		{Id: 2, Name: "Терминатор", Release: 1984, Grade: 8.0},
		{Id: 3, Name: "Побег из Шоушенка", Release: 1994, Grade: 9.1},
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Фильмы",
		})
		c.HTML(http.StatusOK, "Users.html", gin.H{
			"Users": users,
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
