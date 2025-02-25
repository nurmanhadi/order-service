package main

import (
	"context"
	"order-service/config"
	"order-service/event"
)

func main() {
	ctx := context.Background()
	viper := config.NewViper()
	log := config.NewLogrus()
	chProducer := config.NewRabbitmqProducer(viper)
	event.NewOrderEvent(log, chProducer, ctx, viper)
}
