package app

type consumerUseCase struct {
	ent consumerEntity
}

func (cuc *consumerUseCase) saveConsumerData(data interface{}) (interface{}, error) {
	return cuc.ent.Save(data)
}

func newConsumerUseCase() consumerUseCase {
	return consumerUseCase{
		ent: newConsumerEntity(),
	}
}
