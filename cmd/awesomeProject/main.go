package main

import (
	"awesomeProject/internal/app/config"
	"awesomeProject/internal/pkg/app"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Films
// @version 1.0
// @description The best place to watch movies

// @contact.name API Support
// @contact.url https://vk.com/id175719571
// @contact.email stebunov2002@mail.ru

// @license.name AS IS (NO WARRANTY)

// @host 0.0.0.0
// @schemes https http
// @BasePath /

func main() {
	ctx := context.Background()
	log.Println("app start")
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant init config")

		os.Exit(2)
	}

	ctx = config.WrapContext(ctx, cfg)

	fmt.Println(cfg)
	application, err := app.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t create application")

		os.Exit(2)
	}

	err = application.Run(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t run application")

		os.Exit(2)
	}
	log.Println("app terminated")
}
