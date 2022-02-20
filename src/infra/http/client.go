package http

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type client struct {
	url  string
	auth struct {
		url         *string
		credentials interface{}
		token       *string
	}
	restCli *resty.Client
	headers map[string]string
}

func (c *client) shouldAuth() bool {
	return c.hasAuth() && c.auth.token == nil
}

func (c *client) hasAuth() bool {
	return c.auth.url != nil && c.auth.credentials != nil
}

func (c *client) authenticate() (token string, err error) {
	var auth authResponse
	if err = c.post(*c.auth.url, c.auth.credentials, &auth, nil); err == nil {
		token = fmt.Sprintf("%s %s", auth.Type, auth.Token)
	}

	return token, err
}

func (c *client) post(url string, body interface{}, res interface{}, addHeads *map[string]string) (err error) {
	var bys []byte

	if bys, err = json.Marshal(body); err == nil {
		var resErr *httpError

		heads := c.headers
		if addHeads != nil {
			for k, v := range *addHeads {
				heads[k] = v
			}
		}

		_, err = c.restCli.R().
			SetHeaders(heads).
			SetBody(bys).
			SetResult(&res).
			SetError(&resErr).
			Post(url)

		if err == nil {
			if resErr != nil {
				err = errors.New(resErr.Message)
			}
		}
	}

	fmt.Println("res", res)
	fmt.Println("err", err)

	return err
}

func (c *client) Post(body interface{}, path string) (res interface{}, err error) {
	addHead := make(map[string]string)

	// TODO: save token and re-use it until expirition
	if c.shouldAuth() {
		var tok string
		if tok, err = c.authenticate(); err == nil {
			addHead["Authorization"] = tok
		}
	}

	err = c.post(c.url+path, body, res, &addHead)

	return res, err
}

func NewRestClient(baseUrl string, authUrl *string, authCre interface{}) RestClient {
	heads := make(map[string]string)
	heads["Content-Type"] = "application/json"

	return &client{
		url:     baseUrl,
		headers: heads,
		restCli: resty.New(),
		auth: struct {
			url         *string
			credentials interface{}
			token       *string
		}{
			authUrl,
			authCre,
			nil,
		},
	}
}
