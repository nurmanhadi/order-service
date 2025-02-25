package event

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

type IOrderEvent interface {
	PublishPaymentEvent(data []byte) error
}
type orderEvent struct {
	log        *logrus.Logger
	chProducer *amqp091.Channel
	ctx        context.Context
	viper      *viper.Viper
}

func NewOrderEvent(log *logrus.Logger, chProducer *amqp091.Channel, ctx context.Context, viper *viper.Viper) IOrderEvent {
	return &orderEvent{log: log, chProducer: chProducer, ctx: ctx, viper: viper}
}
func (e *orderEvent) PublishPaymentEvent(data []byte) error {
	message := amqp091.Publishing{
		Body: data,
	}
	exchange := e.viper.GetString("rabbitmq.exchange.name")
	routeKey := e.viper.GetString("rabbitmq.queue.payment.route")
	err := e.chProducer.PublishWithContext(e.ctx, exchange, routeKey, false, false, message)
	if err != nil {
		e.log.WithError(err).Error("failed publish message to queue payment")
		return err
	}
	return nil
}
