package handler

import (
	"strconv"

	rajaongkirrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/rajaongkir_request"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initRajaongkirGroup(api *fiber.App) {
	rajaongkirGroup := api.Group("/api/rajaongkir")

	rajaongkirGroup.Get("/provinsi", h.handlerRajaOngkirProvince)
	rajaongkirGroup.Get("/city/:id", h.handlerRajaOngkirCity)
	rajaongkirGroup.Post("/cost", h.handlerRajaOngkirCost)
}

// @Summary Get provinces from RajaOngkir
// @Description Get list of provinces from RajaOngkir
// @Tags RajaOngkir
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 500 {object} responses.ErrorMessage
// @Router /rajaongkir/provinsi [get]
func (h *Handler) handlerRajaOngkirProvince(c *fiber.Ctx) error {
	province, err := h.services.RajaOngkir.GetProvince()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve province data",
		StatusCode: fiber.StatusOK,
		Data:       province,
	})
}

// @Summary Get cities by province ID from RajaOngkir
// @Description Get list of cities by province ID from RajaOngkir
// @Tags RajaOngkir
// @Param id path integer true "Province ID"
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /rajaongkir/city/{id} [get]
func (h *Handler) handlerRajaOngkirCity(c *fiber.Ctx) error {
	id := c.Params("id")

	provinceId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	city, err := h.services.RajaOngkir.GetCity(provinceId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve city data",
		StatusCode: fiber.StatusOK,
		Data:       city,
	})
}

// @Summary Get shipping cost from RajaOngkir
// @Description Get shipping cost from RajaOngkir
// @Tags RajaOngkir
// @Accept json
// @Produce json
// @Param request body rajaongkirrequest.OngkosRequest true "Request body"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /rajaongkir/cost [post]
func (h *Handler) handlerRajaOngkirCost(c *fiber.Ctx) error {
	var request rajaongkirrequest.OngkosRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Invalid request payload",
			},
		)
	}

	cost, err := h.services.RajaOngkir.GetCost(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(cost)
}
