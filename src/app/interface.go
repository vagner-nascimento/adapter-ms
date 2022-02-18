package app

type ConsumerHandler interface {
	Save(data []byte) (interface{}, error)
}
