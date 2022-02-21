package app

import (
	"adapter/src/infra"
)

type consumerEntity struct {
	repo infra.ConsumerDataHandler
}

func (ce *consumerEntity) Save(data interface{}) (interface{}, error) {
	return ce.repo.Save(data)
}

func newConsumerEntity() consumerEntity {
	return consumerEntity{
		repo: infra.NewCosumerRepository(),
	}
}
