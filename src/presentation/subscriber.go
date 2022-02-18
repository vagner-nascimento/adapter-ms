package presentation

import (
	"adapter/src/infra"
	"fmt"
	"os"
)

func SubscribeConsumers() <-chan error {
	c := consumer{
		name:  os.Getenv("AMQP_CONSUMER"),
		topic: os.Getenv("AMQP_TOPIC"),
		handler: func(d []byte) {
			// TODO: implement handler (app Adapter)
			fmt.Println("consumer data", string(d))
		},
	}

	sub := infra.NewSubscriberRepository()

	return sub.SubscribeAll(&c)
}
