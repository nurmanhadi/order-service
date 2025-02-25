package main

import (
	"context"
	"fmt"
	"order-service/config"
)

func main() {
	viper := config.NewViper()
	db := config.NewMysql(viper)
	validation := config.NewValidator()
	app := config.NewFiber(viper)
	log := config.NewLogrus()
	ctx := context.Background()
	chProducer := config.NewRabbitmqProducer(viper)
	config.Bootstrap(&config.BootstrapConfig{
		DB:         db,
		App:        app,
		Ctx:        ctx,
		Log:        log,
		Validation: validation,
		Viper:      viper,
		ChProducer: chProducer,
	})
	webHost := viper.GetString("server.host")
	webPort := viper.GetInt("server.port")
	err := app.Listen(fmt.Sprintf("%s:%d", webHost, webPort))
	if err != nil {
		log.Fatalf("failed to start web server: %v", err)
	}
}
