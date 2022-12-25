package main

import (
	"awesomeProject/internal/pkg/app"
	"context"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Films
// @version 1.0
// @description The best place to watch movies

// @contact.name API Support
// @contact.url https://vk.com/id175719571
// @contact.email stebunov2002@mail.ru

// @host localhost:8080
// @schemes http https
// @BasePath /

func main() {
	log.Println("application start")

	ctx := context.Background()

	application, err := app.New(ctx)
	if err != nil {
		log.Printf("can`t create application: %s", err)
		os.Exit(2)
	}

	err = application.Run()
	if err != nil {
		log.Printf("can`t run application: %s", err)
		os.Exit(2)
	}

	log.Println("application terminated")
}
