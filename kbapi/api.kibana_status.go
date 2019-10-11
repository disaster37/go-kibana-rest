package kbapi

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaStatus = "/api/status" // Base URL to access on Kibana status
)

type KibanaStatus map[string]interface{}

type KibanaStatusGet func() (KibanaStatus, error)

// newKibanaStatusGetFunc permit to get the kibana status and some usefull information
func newKibanaStatusGetFunc(c *resty.Client) KibanaStatusGet {
	return func() (KibanaStatus, error) {
		resp, err := c.R().Get(basePathKibanaStatus)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			if resp.StatusCode() == 404 {
				return nil, nil
			} else {
				return nil, NewAPIError(resp.StatusCode(), resp.Status())
			}
		}
		kibanaStatus := make(KibanaStatus)
		err = json.Unmarshal(resp.Body(), &kibanaStatus)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaStatus: ", kibanaStatus)

		return kibanaStatus, nil
	}
}
