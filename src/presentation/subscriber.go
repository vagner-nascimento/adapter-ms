package presentation

import (
	"adapter/src/app"
	"adapter/src/infra"
	"fmt"
	"os"
)

func SubscribeConsumers() <-chan error {
	c := consumer{
		name:  os.Getenv("AMQP_CONSUMER"),
		topic: os.Getenv("AMQP_TOPIC"),
		handler: func(d []byte) {
			fmt.Println("consumer data", string(d))

			consAdp := app.NewConsumerAdapter()
			consAdp.Save(d)
		},
	}

	sub := infra.NewSubscriberRepository()

	return sub.SubscribeAll(&c)
}
