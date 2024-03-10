package handler

import (
	_ "github.com/alifdwt/haiwan-go/docs"
	"github.com/alifdwt/haiwan-go/internal/service"
	"github.com/alifdwt/haiwan-go/pkg/auth"
	"github.com/alifdwt/haiwan-go/pkg/cloudinary"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type Handler struct {
	services     *service.Service
	cloudinary   cloudinary.MyCloudinary
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Service, cloudinary cloudinary.MyCloudinary, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		cloudinary:   cloudinary,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init() *fiber.App {
	router := fiber.New()

	router.Use(logger.New(
		logger.Config{
			Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "UTC",
		},
	))
	router.Use(cors.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	router.Get("/docs/*", swagger.HandlerDefault)

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *fiber.App) {
	h.initAuthGroup(router)
	h.initUserGroup(router)
	h.initCategoryGroup(router)
	h.initProductGroup(router)
	h.initCartGroup(router)
}
