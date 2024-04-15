package handler

import (
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/slider"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initSliderGroup(api *fiber.App) {
	slider := api.Group("/api/slider")

	slider.Get("/", h.handlerSliderAll)

	slider.Use(middleware.Protector())
	slider.Get("/:id", h.handlerSliderById)
	slider.Post("/create", h.handlerCreateSlider)
	slider.Put("/update/:id", h.handlerSliderUpdate)
	slider.Delete("/delete/:id", h.handlerSliderDelete)
}

// @Summary Get all sliders
// @Description Get all sliders
// @Tags Slider
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /slider [get]
func (h *Handler) handlerSliderAll(c *fiber.Ctx) error {
	res, err := h.services.Slider.GetSliderAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve all sliders",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get slider by ID
// @Description Get slider by ID
// @Tags Slider
// @Param id path integer true "Slider ID"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /slider/{id} [get]
func (h *Handler) handlerSliderById(c *fiber.Ctx) error {
	sliderIdStr := c.Params("id")

	sliderId, err := strconv.Atoi(sliderIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Slider.GetSliderById(sliderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully retrieve slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create a new slider
// @Description Create a new slider
// @Tags Slider
// @Accept multipart/form-data
// @Param name formData string true "Slider Name"
// @Param file formData file true "Slider Image File"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /slider/create [post]
func (h *Handler) handlerCreateSlider(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")
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
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "Failed to upload to Cloudinary: " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	createReq := &slider.CreateSliderRequest{
		Nama:     name,
		FilePath: imageUrl,
	}

	res, err := h.services.Slider.CreateSlider(*createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully create slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update slider by ID
// @Description Update slider by ID
// @Tags Slider
// @Accept multipart/form-data
// @Param id path integer true "Slider ID"
// @Param name formData string true "Slider Name"
// @Param file formData file true "Slider Image File"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /slider/update/{id} [put]
func (h *Handler) handlerSliderUpdate(c *fiber.Ctx) error {
	sliderStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	name := c.FormValue("name")

	file, err := c.FormFile("file")
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
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    "Failed to upload to Cloudinary: " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq := &slider.UpdateSliderRequest{
		Nama:     name,
		FilePath: imageUrl,
	}

	res, err := h.services.Slider.UpdateSliderById(sliderId, *updateReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully update slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete slider by ID
// @Description Delete slider by ID
// @Tags Slider
// @Param id path integer true "Slider ID"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /slider/delete/{id} [delete]
func (h *Handler) handlerSliderDelete(c *fiber.Ctx) error {
	sliderIdStr := c.Params("id")

	sliderId, err := strconv.Atoi(sliderIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Slider.DeleteSliderById(sliderId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully delete slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
