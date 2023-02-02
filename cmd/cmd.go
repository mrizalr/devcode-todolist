package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/devcode-todolist/config"
	"github.com/mrizalr/devcode-todolist/handlers"
	"github.com/mrizalr/devcode-todolist/models"
	"github.com/mrizalr/devcode-todolist/pkg/db"
)

func StartApplication() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config : %v", err)
	}

	db, err := db.NewMysqlConn()
	if err != nil {
		log.Fatalf("Error connect database : %v", err)
	}
	db.AutoMigrate(&models.Activity{}, &models.Todo{})

	app := fiber.New()
	handlers.MapHandler(app, db)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
