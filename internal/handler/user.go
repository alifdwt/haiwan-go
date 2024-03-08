package handler

import (
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initUserGroup(api *fiber.App) {
	user := api.Group("/api/user")

	user.Get("/", h.handlerUserAll)
	user.Get("/:id", h.handlerUserById)

	user.Use(middleware.Protector())
	user.Post("/create", h.handlerCreateUser)
	user.Put("/update/:id", h.handlerUpdateUserById)
	user.Delete("/delete/:id", h.handlerDeleteUserById)
}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user [get]
func (h *Handler) handlerUserAll(c *fiber.Ctx) error {
	res, err := h.services.User.GetUserAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully get all user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags User
// @Param id path integer true "User ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user/{id} [get]
func (h *Handler) handlerUserById(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.User.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get user by ID",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body auth.RegisterRequest true "User details"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user/create [post]
func (h *Handler) handlerCreateUser(c *fiber.Ctx) error {
	var body auth.RegisterRequest

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

	res, err := h.services.User.CreateUser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully create user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update user by ID
// @Description Update user by ID
// @Tags User
// @Param id path integer true "User ID"
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body user.UpdateUserRequest true "User details"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user/update/{id} [put]
func (h *Handler) handlerUpdateUserById(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	var body user.UpdateUserRequest

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

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

	res, err := h.services.User.UpdateUserById(userId, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully update user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags User
// @Param id path integer true "User ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user/delete/{id} [delete]
func (h *Handler) handlerDeleteUserById(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.User.DeleteUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully delete user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
