package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecanutto/isaac-structure/src/handlers"
	"github.com/alecanutto/isaac-structure/src/repositories"
	"github.com/alecanutto/isaac-structure/src/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"babar.db.elephantsql.com", 5432, "xkpqgdbw", "S1tbcBoQOF5HkHeLEZmRw4spLU3eU4Vv", "xkpqgdbw")

	dialector := postgres.Open(psqlInfo)

	db, err := gorm.Open(dialector)

	if err != nil {
		log.Fatalf("Falha na conexão com db. Err=%v", err.Error())
	}

	repoContainer := repositories.GetRepositories(db)

	servContainer := services.GetServices(repoContainer)

	handlers.NewHandlerContainer(server, servContainer)

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("OK")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("Servidor não iniciado. Err=%v", err.Error())
	}
}
