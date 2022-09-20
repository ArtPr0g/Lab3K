package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

type Viewdata struct {
	Title string
	Users []Users
}
type Users struct {
	Name    string
	Release int
	Grade   float64
}

func StartServer() {
	log.Println("Server start up")

	data := Viewdata{
		Title: "List of films",
		Users: []Users{
			Users{Name: "Крепкий орешек", Release: 1988, Grade: 8.0},
			Users{Name: "Терминатор", Release: 1984, Grade: 8.0},
			Users{Name: "Побег из Шоушенка", Release: 1994, Grade: 9.1},
		},
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

	r.Static("/image", "./resources")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("templates/Users.html")
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
