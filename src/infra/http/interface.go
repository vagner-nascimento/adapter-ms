package http

type RestClient interface {
	Post(body interface{}, response interface{}, path string) error
}
