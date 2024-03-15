package handler

import (
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/review"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initReviewGroup(api *fiber.App) {
	reviewGroup := api.Group("/api/review")

	reviewGroup.Get("/", h.handlerGetReviewAll)
	reviewGroup.Get("/:id", h.handlerGetReviewById)

	reviewGroup.Use(middleware.Protector())
	reviewGroup.Post("/create", h.handlerCreateReview)
}

// @Summary Get all reviews
// @Description Get all reviews
// @Tags Review
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 500 {object} responses.ErrorMessage
// @Router /review [get]
func (h *Handler) handlerGetReviewAll(c *fiber.Ctx) error {
	reviews, err := h.services.Review.GetReviewAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve all reviews",
		StatusCode: fiber.StatusOK,
		Data:       reviews,
	})
}

// @Summary Get review by ID
// @Description Get review by ID
// @Tags Review
// @Param id path integer true "Review ID"
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /review/{id} [get]
func (h *Handler) handlerGetReviewById(c *fiber.Ctx) error {
	reviewId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	review, err := h.services.Review.GetReviewById(reviewId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve review by Id",
		StatusCode: fiber.StatusOK,
		Data:       review,
	})
}

// @Summary Create review
// @Description Create review
// @Tags Review
// @Param Authorization header string true "JWT token"
// @Param product_id path integer true "Product ID"
// @Accept json
// @Produce json
// @Param request body review.CreateReviewRequest true "Review data"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /review/create/{product_id} [post]
func (h *Handler) handlerCreateReview(c *fiber.Ctx) error {
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

	productId, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body review.CreateReviewRequest
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

	res, err := h.services.Review.CreateReview(productId, userId, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully created review",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
