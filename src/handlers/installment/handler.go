package installment

import (
	"net/http"

	"github.com/alecanutto/isaac-structure/src/interfaces"
	"github.com/alecanutto/isaac-structure/src/services"
	"github.com/alecanutto/isaac-structure/src/structs"
	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	router             fiber.Router
	installmentService interfaces.InstallmentService
}

func NewInstallmentHandler(router fiber.Router, servContainer services.ServiceContainer) InstallmentHandler {
	return InstallmentHandler{
		router:             router,
		installmentService: servContainer.InstallmentService,
	}
}

func (ih InstallmentHandler) SetRoutes() {
	groupRouter := ih.router.Group("/installments")
	groupRouter.Post("", ih.CreateInstallment)
	groupRouter.Get("/last_item", ih.GetLastItem)
}

func (ih InstallmentHandler) CreateInstallment(ctx *fiber.Ctx) error {
	var body structs.Installments

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag:  "BAD_REQUEST",
		})
	}

	if err := ih.installmentService.Create(body); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(structs.Response{
			Data: err.Error(),
			Tag:  "INTERNAL_SERVER_ERROR",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(structs.Response{
		Data: "Installment criado com sucesso",
	})
}

func (ih InstallmentHandler) GetLastItem(ctx *fiber.Ctx) error {
	item, err := ih.installmentService.GetLastItem()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(structs.Response{
			Data: err.Error(),
			Tag:  "INTERNAL_SERVER_ERROR",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(item)
}
