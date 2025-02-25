package usecase

import (
	"order-service/event"
	"order-service/internal/entity"
	"order-service/internal/model"
	"order-service/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type IOrderUsecase interface {
	AddOrder(request *model.OrderRequestData) error
}
type orderUsecase struct {
	orderRepository repository.IOrderRepository
	validation      *validator.Validate
	log             *logrus.Logger
	orderEvent      event.IOrderEvent
}

func NewOrderUsecase(
	orderRepository repository.IOrderRepository,
	validation *validator.Validate,
	log *logrus.Logger,
	orderEvent event.IOrderEvent,
) IOrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
		validation:      validation,
		log:             log,
		orderEvent:      orderEvent,
	}
}
func (s *orderUsecase) AddOrder(request *model.OrderRequestData) error {
	err := s.validation.Struct(request)
	if err != nil {
		s.log.WithField("error", err).Warn("failed validation request")
		return err
	}
	order := &entity.Order{
		Id:        request.OrderId,
		UserId:    request.UserId,
		ProductId: request.ProductId,
		Price:     request.Price,
		Quantity:  request.Quantity,
		Status:    entity.STATUS_PENDING,
	}
	err = s.orderRepository.Add(order)
	if err != nil {
		s.log.WithField("error", err).Warn("failed add order from database")
		return err
	}
	orderJSON, err := json.Marshal(order)
	if err != nil {
		s.log.WithField("error", err).Warn("failed to serialize order to json")
		return err
	}
	go func() {
		err = s.orderEvent.PublishPaymentEvent(orderJSON)
		if err != nil {
			s.log.WithField("error", err).Warn("failed to publish order event")
		}
	}()
	return nil
}
