package handler

import (
	"fmt"
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initProductGroup(api *fiber.App) {
	product := api.Group("/api/product")

	product.Get("/hello", h.handlerProductHello)
	product.Get("/", h.handlerProductAll)
	product.Get("/slug/:slug", h.handlerProductBySlug)

	product.Use(middleware.Protector())
	product.Get("/:id", h.handlerProductById)
	product.Post("/create", h.handlerProductCreate)
	product.Put("/update/:id", h.handlerProductUpdate)
	product.Delete("/delete/:id", h.handlerProductDelete)
}

// @Summary Greet the user for products
// @Description Return a greeting message for products
// @Tags Product
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /product/hello [get]
func (h *Handler) handlerProductHello(c *fiber.Ctx) error {
	return c.SendString("This is a handler Product")
}

// @Summary Get all products
// @Description Retrieve all products
// @Tags Product
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /product [get]
func (h *Handler) handlerProductAll(c *fiber.Ctx) error {
	res, err := h.services.Product.GetProductAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve all products",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get product by ID
// @Description Retrieve a product by its ID
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /product/{id} [get]
func (h *Handler) handlerProductById(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.GetProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully retrieve product data by id",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get product by slug
// @Description Retrieve a product by its slug
// @Tags Product
// @Produce json
// @Param slug path string true "Product Slug"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /product/slug/{slug} [get]
func (h *Handler) handlerProductBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.services.Product.GetProductBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully retrieve product data by slug",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create product
// @Description Create a new product
// @Tags Product
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Product Name"
// @Param category formData string true "Category ID"
// @Param description formData string true "Product Description"
// @Param brand formData string true "Product Brand"
// @Param price formData integer true "Product Price"
// @Param countInStock formData integer true "Product Count In Stock"
// @Param weight formData integer true "Product Weight"
// @Param rating formData integer true "Product Rating"
// @Param image_product formData file true "Product Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /product/create [post]
func (h *Handler) handlerProductCreate(c *fiber.Ctx) error {
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	countInStock, err := strconv.Atoi(c.FormValue("countInStock"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	createReq := product.CreateProductRequest{
		Name:         c.FormValue("name"),
		CategoryID:   c.FormValue("category"),
		Description:  c.FormValue("description"),
		Brand:        c.FormValue("brand"),
		Price:        price,
		CountInStock: countInStock,
		Weight:       weight,
		Rating:       &rating,
	}

	file, err := c.FormFile("image_product")
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
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s-1", slug.Make(createReq.Name))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	createReq.FilePath = imageUrl

	if err := createReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.CreateProduct(&createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully creating product",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update product
// @Description Update an existing product
// @Tags Product
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param name formData string true "Product Name"
// @Param category formData string true "Category ID"
// @Param description formData string true "Product Description"
// @Param brand formData string true "Product Brand"
// @Param price formData integer true "Product Price"
// @Param countInStock formData integer true "Product Count In Stock"
// @Param weight formData integer true "Product Weight"
// @Param rating formData integer true "Product Rating"
// @Param image_product formData file true "Product Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /product/update/{id} [put]
func (h *Handler) handlerProductUpdate(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	countInStock, err := strconv.Atoi(c.FormValue("countInStock"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	updateReq := product.UpdateProductRequest{
		Name:         c.FormValue("name"),
		CategoryID:   c.FormValue("category"),
		Description:  c.FormValue("description"),
		Brand:        c.FormValue("brand"),
		Price:        price,
		CountInStock: countInStock,
		Weight:       weight,
		Rating:       rating,
	}

	file, err := c.FormFile("image_product")
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
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s-1", slug.Make(updateReq.Name))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq.FilePath = imageUrl

	if err := updateReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.UpdateProduct(productId, &updateReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully update product",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete product
// @Description Delete an existing product
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /product/delete/{id} [delete]
func (h *Handler) handlerProductDelete(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	producId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.DeleteProduct(producId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully delete product",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
