package router

import (
	"order-service/internal/delivery/rest"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App          *fiber.App
	OrderHandler rest.IOrderHandler
}

func (r *RouteConfig) Router() {
	api := r.App.Group("/api/v1")

	order := api.Group("/orders")
	order.Post("/", r.OrderHandler.AddOrder)
}
