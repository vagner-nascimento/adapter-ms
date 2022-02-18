package infra

type SubscribeDataHandler interface {
	SubscribeAll(cs ...Consumer) <-chan error
}

type ConsumerDataHandler interface {
	Save(data []byte) (interface{}, error)
}

type Consumer interface {
	Name() string
	Tpoic() string
	Handler() func([]byte)
}
