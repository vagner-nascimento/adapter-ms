package presentation

import (
	"adapter/src/app"
	"adapter/src/infra"
	"encoding/json"
	"fmt"
	"os"
)

func SubscribeConsumers() <-chan error {
	c := consumer{
		name:  os.Getenv("AMQP_CONSUMER"),
		topic: os.Getenv("AMQP_TOPIC"),
		handler: func(d []byte) {
			fmt.Println("consumer data", string(d))

			var obj interface{}
			if err := json.Unmarshal(d, &obj); err == nil {
				consAdp := app.NewConsumerAdapter()

				consAdp.Save(obj)
			} else {
				fmt.Println("consumer parse error", err)
			}
		},
	}

	sub := infra.NewSubscriberRepository()

	return sub.SubscribeAll(&c)
}
