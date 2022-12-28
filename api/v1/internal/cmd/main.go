package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kraikub/mail-service/api/v1/internal/config"
	"github.com/kraikub/mail-service/api/v1/internal/controllers"
	"github.com/kraikub/mail-service/api/v1/internal/usecases"
	"github.com/kraikub/mail-service/servers"
)

func main() {

	config, err := config.GetRuntimeConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	// ctx, _ := context.WithCancel(context.Background())
	mail := usecases.CreateMailUseCase(
		config.Smtp.Host,
		strconv.Itoa(config.Smtp.Port),
		config.Smtp.ServiceEmail,
		config.Smtp.ServiceEmailPassword,
	)

	kraikub := servers.NewKraikubServer(config.Server.Name, config.Server.Port)
	controllers.AssignRouter(
		kraikub.Router(),
		mail,
	)

	// No need any go routines
	kraikub.StartWithGraceFullShutdown(func(cancel context.CancelFunc) {
		// if err = mongoClient.Disconnect(context.TODO()); err != nil {
		// should not panic
		// 	log.Fatal(err)
		// }
		cancel()
	})
}
