package config

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func NewRabbitmqProducer(viper *viper.Viper) *amqp091.Channel {
	user := viper.GetString("rabbitmq.user")
	password := viper.GetString("rabbitmq.password")
	host := viper.GetString("rabbitmq.host")
	port := viper.GetString("rabbitmq.port")
	virtualHost := viper.GetString("rabbitmq.virtual_host")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, virtualHost)

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}
	return ch
}
func NewRabbitmqConsumer(viper *viper.Viper) *amqp091.Channel {
	user := viper.GetString("rabbitmq.user")
	password := viper.GetString("rabbitmq.password")
	host := viper.GetString("rabbitmq.host")
	port := viper.GetString("rabbitmq.port")
	virtualHost := viper.GetString("rabbitmq.virtual_host")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, virtualHost)

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}
	return ch
}
