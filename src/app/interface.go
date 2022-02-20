package app

type ConsumerHandler interface {
	Save(data interface{}) (interface{}, error)
}
