package http

import (
	"order-service/internal/model"
	"order-service/internal/usecase"
	"order-service/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type IOrderHandler interface {
	AddOrder(c *fiber.Ctx) error
}
type orderHandler struct {
	orderUsecase usecase.IOrderUsecase
	log          *logrus.Logger
}

func NewOrderHandler(orderUsecase usecase.IOrderUsecase, log *logrus.Logger) IOrderHandler {
	return &orderHandler{
		orderUsecase: orderUsecase,
		log:          log,
	}
}
func (h *orderHandler) AddOrder(c *fiber.Ctx) error {
	request := new(model.OrderRequestData)
	err := c.BodyParser(&request)
	if err != nil {
		h.log.WithField("error", err).Warn("failed parse request")
		return response.ErrorR(c, 400, "failed parse request")
	}
	err = h.orderUsecase.AddOrder(request)
	if err != nil {
		return response.Error(c, err)
	}

	h.log.WithField("data", request).Info("success add order")
	return response.Success(c, 200, nil)
}
