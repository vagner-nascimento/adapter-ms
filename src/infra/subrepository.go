package infra

import (
	"adapter/src/infra/amqp"
	"fmt"
)

type subRepository struct {
	sub amqp.Subscriber
}

func (sr *subRepository) SubscribeAll(cs ...Consumer) <-chan error {
	errCh := make(chan error)
	var err error

	for _, c := range cs {
		cons := c.Name()
		tp := c.Tpoic()

		if err = sr.sub.Subscribe(cons, tp, c.Handler()); err != nil {
			go func() {
				fmt.Println("error", err)
				errCh <- err
			}()

			break
		}

		msg := fmt.Sprintf("subscribed %s into %s", cons, tp)
		fmt.Println(msg)
	}

	if err == nil {
		amqClosed := amqp.ListenAmqpCloseConn()

		go func() {
			for err := range amqClosed {
				if err != nil {
					errCh <- err
				}
			}
		}()
	}

	return errCh
}

func NewSubscriberRepository() SubscribeDataHandler {
	return &subRepository{
		sub: amqp.NewSubscriber(),
	}
}
