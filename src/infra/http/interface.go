package http

type RestClient interface {
	Post(body interface{}, path string) (interface{}, error)
}
