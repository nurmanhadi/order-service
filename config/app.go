package config

import (
	"context"
	"database/sql"
	"order-service/event"
	"order-service/internal/delivery/http"
	"order-service/internal/delivery/http/router"
	"order-service/internal/repository"
	"order-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB         *sql.DB
	App        *fiber.App
	Ctx        context.Context
	Log        *logrus.Logger
	Validation *validator.Validate
	Viper      *viper.Viper
	ChProducer *amqp091.Channel
}

func Bootstrap(config *BootstrapConfig) {
	// event
	orderEvent := event.NewOrderEvent(config.Log, config.ChProducer, config.Ctx, config.Viper)

	// repositories
	orderRepository := repository.NewOrderRepository(config.DB, config.Log, config.Ctx)

	// usecase
	orderUsecase := usecase.NewOrderUsecase(orderRepository, config.Validation, config.Log, orderEvent)

	// handler
	orderHandler := http.NewOrderHandler(orderUsecase, config.Log)

	routeConfig := &router.RouteConfig{
		App:          config.App,
		OrderHandler: orderHandler,
	}
	routeConfig.Router()
}
