package amqp

type Subscriber interface {
	Subscribe(consu string, topic string, handler func([]byte)) error
}
