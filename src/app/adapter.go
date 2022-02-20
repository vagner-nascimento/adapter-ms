package app

import (
	"adapter/src/infra"
)

type consAdp struct{}

func (c *consAdp) Save(data interface{}) (interface{}, error) {
	repo := infra.NewCosumerRepository()

	return repo.Save(data)
}

func NewConsumerAdapter() ConsumerHandler {
	return &consAdp{}
}
