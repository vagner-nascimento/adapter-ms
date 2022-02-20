package infra

import (
	"adapter/src/infra/http"
	"os"
)

type consRepository struct {
	cli http.RestClient
}

type consAuth struct {
	GranTp string `json:"grant_type"`
	CliId  string `json:"client_id"`
	CliSec string `json:"client_secret"`
	AccId  string `json:"account_id"`
}

func (cr *consRepository) Save(data interface{}) (interface{}, error) {
	return cr.cli.Post(data, "")
}

func NewCosumerRepository() ConsumerDataHandler {
	url := os.Getenv("CONSUMER_URL")
	authUrl := os.Getenv("CONSUMER_AUTH_URL")
	cliId := os.Getenv("CONSUMER_AUTH_ID")
	cliSec := os.Getenv("CONSUMER_AUTH_SCRET")
	accId := os.Getenv("CONSUMER_AUTH_ACC_ID")
	graTp := os.Getenv("CONSUMER_AUTH_GRANT_TP")

	authCre := consAuth{
		GranTp: graTp,
		CliId:  cliId,
		CliSec: cliSec,
		AccId:  accId,
	}

	return &consRepository{
		cli: http.NewRestClient(url, &authUrl, authCre),
	}
}
