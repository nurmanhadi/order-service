package router

import (
	"order-service/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App          *fiber.App
	OrderHandler http.IOrderHandler
}

func (r *RouteConfig) Router() {
	api := r.App.Group("/api/v1")

	order := api.Group("/orders")
	order.Post("/", r.OrderHandler.AddOrder)
}
