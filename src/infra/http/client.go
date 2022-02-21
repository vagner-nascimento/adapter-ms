package http

import (
	"adapter/src/utils"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
)

type client struct {
	url  string
	auth struct {
		url         *string
		credentials interface{}
	}
	restCli *resty.Client
	headers map[string]string
}

func (c *client) shouldAuth() bool {
	_, exists := c.headers["Authorization"]
	return c.hasAuth() && !exists
}

func (c *client) hasAuth() bool {
	return c.auth.url != nil && c.auth.credentials != nil
}

func (c *client) authenticate(refresh bool) (err error) {
	if refresh {
		delete(c.headers, "Authorization")
	}

	var auth authResponse
	if err = c.post(*c.auth.url, c.auth.credentials, &auth, nil); err == nil {
		token := fmt.Sprintf("%s %s", auth.Type, auth.Token)
		c.headers["Authorization"] = token

		go func() {
			fmt.Println(fmt.Sprintf("Auth token store for %d seconds", auth.Expires))

			utils.Sleep(int64(auth.Expires))
			delete(c.headers, "Authorization")

			fmt.Println("Auth token deleted")
		}()
	}

	return err
}

func (c *client) post(url string, body interface{}, res interface{}, addHeads *map[string]string) (err error) {
	var bys []byte

	if bys, err = json.Marshal(body); err == nil {
		var resErr *httpError

		heads := make(map[string]string)
		if addHeads != nil {
			utils.FillMap(heads, *addHeads)
		}

		// default headers should be filled after aditionals to avoid overriding
		utils.FillMap(heads, c.headers)

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

	return err
}

func (c *client) Post(body interface{}, res interface{}, path string) (err error) {
	if c.shouldAuth() {
		err = c.authenticate(false)
	}

	if err == nil {
		if err = c.post(c.url+path, body, res, nil); err != nil {
			if strings.ToLower(err.Error()) == "not authorized" && !c.shouldAuth() {
				err = c.authenticate(true)
				err = c.post(c.url+path, body, res, nil)
			}
		}
	}

	return err
}

var singClient struct {
	start  sync.Once
	client *client
}

func NewRestClient(baseUrl string, authUrl *string, authCre interface{}) RestClient {
	heads := make(map[string]string)
	heads["Content-Type"] = "application/json"

	singClient.start.Do(func() {
		singClient.client = &client{
			url:     baseUrl,
			headers: heads,
			restCli: resty.New(),
			auth: struct {
				url         *string
				credentials interface{}
			}{
				authUrl,
				authCre,
			},
		}
	})

	return singClient.client
}
