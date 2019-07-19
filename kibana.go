package kibana

import (
	"crypto/tls"
	"github.com/disaster37/go-kibana-rest/kbapi"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	Address          string
	Username         string
	Password         string
	DisableVerifySSL bool
}

type Client struct {
	*kbapi.API
	Client *resty.Client
}

func NewDefaultClient() (*Client, error) {
	return NewClient(Config{})
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.Address == "" {
		cfg.Address = "http://localhost:5602"
	}

	restyClient := resty.New().
		SetHostURL(cfg.Address).
		SetBasicAuth(cfg.Username, cfg.Password).
		SetHeader("kbn-xsrf", "true")

	client := &Client{
		Client: restyClient,
		API:    kbapi.New(restyClient),
	}

	if cfg.DisableVerifySSL == true {
		client.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	return client, nil

}
