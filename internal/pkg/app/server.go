package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Users struct {
	Name    string `json:"name"`
	Release int    `json:"release"`
	Grade   int    `json:"grade"`
}

func (a *Application) StartServer() {
	log.Println("Server start up")

	var users = []Users{
		{Name: "Железный человек", Release: 2008, Grade: 7},
		{Name: "Терминатор", Release: 1984, Grade: 8},
		{Name: "Трансформеры", Release: 2007, Grade: 8},
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string
		if id != "" {
			log.Printf("id received %s\n", id)
			intID, err := strconv.Atoi(id)
			if err != nil {
				log.Printf("can't convert id %v", err)
				c.Error(err)
				return
			}

			users, err := a.repo.GetKinoByID(uint(intID))
			if err != nil {
				log.Printf("can't get promo by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"Name":    users.Name,
				"Release": users.Release,
				"grade":   users.Grade,
			})
			return
		}
		create := c.Query("create")
		if create != "" {
			log.Printf("create received %s\n", create)
			createBool, err := strconv.ParseBool(create)
			if err != nil {
				log.Printf("can't convert create %v", err)
				c.Error(err)
				return
			}

			if createBool {
				a.repo.NewRandRecords()
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "create not true",
			})

			return

			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		}
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Список фильмов",
		})
		c.HTML(http.StatusOK, "users.html", gin.H{
			"Users": users,
		})
	})
	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
