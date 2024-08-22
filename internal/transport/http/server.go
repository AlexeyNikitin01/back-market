package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"task5/internal/config"
	usecase "task5/internal/domain"
)

type Server struct {
	App *fiber.App
}

// NewServer TODO закончить инициализацию сервера обработчиками
func NewServer(cfg *config.Config, u *usecase.Usecase) *Server {
	instance := Server{
		App: fiber.New(fiber.Config{ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return nil
		}}),
	}
	h := newHandler(cfg, u)
	instance.App.Use(
		cors.New(cors.Config{
			AllowOrigins:     "https://example.com, https://editor.swagger.io/",
			AllowMethods:     "GET,POST,OPTIONS,DELETE,PUT",
			AllowHeaders:     "origin, x-requested-with, content-type, authorization",
			AllowCredentials: true,
		}),
		recover.New(),
		contextualLoggerMiddleware,
		errorMiddleware,
		httpRequestLoggerMiddleware,
		h.authMiddleware,
		h.permissionMiddleware,
	)
	base := instance.App.Group("/v1")
	base.Get("/user", h.GetUser)
	base.Get("/ping", h.Ping)
	base.Post("/token", h.Token)
	base.Post("/register", h.Register)
	base.Post("/refresh", h.Refresh)

	base.Get("/product/:id", h.GetProduct)
	base.Get("/products", h.GetProductsByQuery)
	base.Post("/products", h.CreateProducts)
	base.Put("/products", h.UpdateProduct)
	base.Delete("/products/:id", h.DeleteProductById)

	base.Get("/cart", h.GetCart)
	base.Put("/cart", h.UpdateCart)

	base.Get("/order/:id", h.GetOrder)
	base.Get("/orders", h.GetOrders)
	base.Post("/order", h.CreateOrder)
	base.Delete("/order/:id", h.DeleteOrder)

	base.Post("/pay", h.Pay)

	return &instance
}
