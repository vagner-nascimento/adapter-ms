package app

type consumerAdp struct {
	useCase consumerUseCase
}

func (ca *consumerAdp) Save(data interface{}) (interface{}, error) {
	return ca.useCase.ent.Save(data)
}

func NewConsumerAdapter() ConsumerHandler {
	return &consumerAdp{
		useCase: newConsumerUseCase(),
	}
}
