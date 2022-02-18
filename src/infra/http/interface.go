package http

type RestClient interface {
	Post(body []byte, path string) (interface{}, error)
}
