package handler

import (
	midtransrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/midtrans_request"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initMidtransGroup(api *fiber.App) {
	midtrans := api.Group("/api/midtrans")

	midtrans.Post("/create-transaction", h.handlerCreateTransaction)
}

// @Summary Create transaction
// @Description Create a new transaction
// @Tags Midtrans
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createMidtransRequest body midtransrequest.CreateMidtransRequest true "Request body to create a new transaction"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /midtrans/create-transaction [post]
func (h *Handler) handlerCreateTransaction(c *fiber.Ctx) error {
	var request midtransrequest.CreateMidtransRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Midtrans.CreateTransaction(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully create transaction",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
