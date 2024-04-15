package handler

import (
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initCategoryGroup(api *fiber.App) {
	category := api.Group("/api/category")

	category.Get("/hello", h.handlerHelloCategory)
	category.Get("/", h.handlerCategoryAll)
	category.Get("/slug/:slug", h.handlerCategoryBySlug)

	category.Use(middleware.Protector())
	category.Get("/:id", h.handlerCategoryById)
	category.Post("/create", h.handlerCategoryCreate)
	category.Put("/update/:id", h.handlerCategoryUpdate)
	category.Delete("/delete/:id", h.handlerCategoryDelete)
}

func (h *Handler) handlerHelloCategory(c *fiber.Ctx) error {
	return c.SendString("This is category handler")
}

// @Summary Get all categories
// @Description Get all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /category [get]
func (h *Handler) handlerCategoryAll(c *fiber.Ctx) error {
	res, err := h.services.Category.GetCategoryAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully get categories",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get category by ID
// @Description Get a category by its ID
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /category/{id} [get]
func (h *Handler) handlerCategoryById(c *fiber.Ctx) error {
	categoryIdStr := c.Params("id")

	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Category.GetCategoryById(categoryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Succesfully get category data",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get category by slug
// @Description Get a category by its slug
// @Tags Category
// @Accept json
// @Produce json
// @Param slug path string true "Category Slug"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /category/slug/{slug} [get]
func (h *Handler) handlerCategoryBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.services.Category.GetCategoryBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get category data by slug",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create category
// @Description Create a new category
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param name formData string true "Category Name"
// @Param image_category formData file true "Category Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /category/create [post]
func (h *Handler) handlerCategoryCreate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("image_category")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "File not found: " + err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "Failed to open uploaded file: " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, slug.Make(name))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			responses.ErrorMessage{
				Error:      true,
				Message:    "Failed to upload file to Cloudinary: " + err.Error(),
				StatusCode: fiber.StatusInternalServerError,
			},
		)
	}

	createReq := &category.CreateCategoryRequest{
		Name:     name,
		FilePath: imageUrl,
	}

	res, err := h.services.Category.CreateCategory(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully create category",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update category by ID
// @Description Update a category by its ID
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param name formData string true "Category Name"
// @Param file formData file true "Category Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /category/update/{id} [put]
func (h *Handler) handlerCategoryUpdate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "file not found " + err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	categoryIdStr := c.Params("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "Failed to open uploaded file: " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "Failed to upload file to Cloudinary: " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq := &category.UpdateCategoryRequest{
		Name:     name,
		FilePath: imageUrl,
	}

	res, err := h.services.Category.UpdateCategoryById(categoryId, updateReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully update category",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete category by ID
// @Description Delete a category by its ID
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /category/delete/{id} [delete]
func (h *Handler) handlerCategoryDelete(c *fiber.Ctx) error {
	categoryIdStr := c.Params("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid category ID",
			"error":   true,
		})
	}

	res, err := h.services.Category.DeleteCategoryById(categoryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "successfully delete category",
		"data":       res,
		"statusCode": fiber.StatusOK,
	})
}
