package handler

import (
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/order"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initOrderGroup(api *fiber.App) {
	order := api.Group("/api/order")

	order.Get("/hello", h.handlerHelloOrder)

	order.Use(middleware.Protector())
	order.Get("/", h.handlerGetOrderAll)
	order.Get("/orderbyuser", h.handlerGetOrderByUserId)
	order.Post("/create", h.handlerCreateOrder)
	order.Get("/:id", h.handlerGetOrderById)
}

// @Summary Greet the user for orders
// @Description Return a greeting message for orders
// @Tags Order
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /order/hello [get]
func (h *Handler) handlerHelloOrder(c *fiber.Ctx) error {
	return c.SendString("This is a handler Order")
}

// @Summary Get all orders
// @Description Retrieve all orders
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /order [get]
func (h *Handler) handlerGetOrderAll(c *fiber.Ctx) error {
	res, err := h.services.Order.GetOrderAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get all order data",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get order by ID
// @Description Retrieve an order by its ID
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /order/{id} [get]
func (h *Handler) handlerGetOrderById(c *fiber.Ctx) error {
	orderIdStr := c.Params("id")

	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Order.GetOrderById(orderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get order by its id",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get orders by user ID
// @Description Retrieve orders associated with a specific user
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /order/orderbyuser [get]
func (h *Handler) handlerGetOrderByUserId(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Order.GetOrdersByUser(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get orders by its user id",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createOrderRequest body order.CreateOrderRequest true "Request body to create a new order"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /order/create [post]
func (h *Handler) handlerCreateOrder(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body order.CreateOrderRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Order.CreateOrder(userId, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully create order",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
