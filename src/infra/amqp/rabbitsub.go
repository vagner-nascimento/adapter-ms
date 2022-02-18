package amqp

import (
	stAmq "github.com/streadway/amqp"
)

type rabbitSub struct{}

func (rs *rabbitSub) Subscribe(consu string, topic string, process func([]byte)) (err error) {
	var ch *stAmq.Channel

	if ch, err = getChannel(); err == nil {
		var q stAmq.Queue

		if q, err = ch.QueueDeclare(topic, false, false, false, false, nil); err == nil {
			var msgs <-chan stAmq.Delivery

			if msgs, err = ch.Consume(q.Name, consu, true, false, false, false, nil); err == nil {
				go func() {
					for msg := range msgs {
						process(msg.Body)
					}
				}()
			}
		}
	}

	return err
}

func NewSubscriber() Subscriber {
	return &rabbitSub{}
}
