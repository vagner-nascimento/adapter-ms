package presentation

type consumer struct {
	name    string
	topic   string
	handler func([]byte)
}

func (c *consumer) Name() string {
	return c.name
}

func (c *consumer) Tpoic() string {
	return c.topic
}

func (c *consumer) Handler() func([]byte) {
	return c.handler
}
