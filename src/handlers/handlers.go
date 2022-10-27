package handlers

import (
	"github.com/alecanutto/isaac-structure/src/handlers/installment"
	"github.com/alecanutto/isaac-structure/src/services"
	"github.com/gofiber/fiber/v2"
)

func NewHandlerContainer(router fiber.Router, servContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(router, servContainer).SetRoutes()
}
