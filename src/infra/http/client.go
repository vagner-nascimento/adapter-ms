package http

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type client struct {
	baseUrl string
	authUrl string
	restCli *resty.Client
	headers map[string]string
}

func (c *client) Post(body []byte, path string) (res interface{}, err error) {
	var resErr interface{}
	var resp *resty.Response

	fmt.Println("bys", string(body))

	resp, err = c.restCli.R().
		SetHeaders(c.headers).
		SetBody(body).
		SetResult(&res).
		SetError(&resErr).
		Post(c.baseUrl + path)

	fmt.Println("POST", res, resErr, resp)

	return res, err
}

func NewRestClient(baseUrl string, authUrl string) RestClient {
	heads := make(map[string]string)
	heads["Content-Type"] = "application/json"

	return &client{
		baseUrl: baseUrl,
		authUrl: authUrl,
		headers: heads,
		restCli: resty.New(),
	}
}

/*
approved := make(map[int]string)
	approved[1] = "Mary"
*/
