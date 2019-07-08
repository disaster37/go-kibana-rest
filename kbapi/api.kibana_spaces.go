package kbapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaSpace = "/api/spaces/space" // Base URL to access on Kibana space API
)

// Kibana space object
type KibanaSpace struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description,omitempty"`
	DisabledFeatures []string `json:"disabledFeatures,omitempty"`
	Reserved         bool     `json:"_reserved,omitempty"`
	Initials         string   `json:"initials,omitempty"`
	Color            string   `json:"color,omitempty"`
}

func (k *KibanaSpace) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}

// List of KibanaSpace objects
type KibanaSpaces []KibanaSpace

type KibanaSpaceGet func(id string) (*KibanaSpace, error)
type KibanaSpaceList func() (KibanaSpaces, error)
type KibanaSpaceCreate func(kibanaSpace *KibanaSpace) (*KibanaSpace, error)
type KibanaSpaceDelete func(id string) error
type KibanaSpaceUpdate func(kibanaSpace *KibanaSpace) (*KibanaSpace, error)

// newKibanaSpaceGetFunc permit to get the kibana space with it id
func newKibanaSpaceGetFunc(c *resty.Client) KibanaSpaceGet {
	return func(id string) (*KibanaSpace, error) {

		if id == "" {
			return nil, NewAPIError(600, "You must provide kibana space ID")
		}
		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s", basePathKibanaSpace, id)
		resp, err := c.R().Get(path)
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
		kibanaSpace := &KibanaSpace{}
		err = json.Unmarshal(resp.Body(), kibanaSpace)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaSpace: ", kibanaSpace)

		return kibanaSpace, nil
	}

}

// newKibanaSpaceListFunc permit to get all Kibana space
func newKibanaSpaceListFunc(c *resty.Client) KibanaSpaceList {
	return func() (KibanaSpaces, error) {

		resp, err := c.R().Get(basePathKibanaSpace)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		kibanaSpaces := make(KibanaSpaces, 0, 1)
		err = json.Unmarshal(resp.Body(), &kibanaSpaces)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaSpaces: ", kibanaSpaces)

		return kibanaSpaces, nil
	}

}

// newKibanaSpaceCreateFunc permit to create new Kibana space
func newKibanaSpaceCreateFunc(c *resty.Client) KibanaSpaceCreate {
	return func(kibanaSpace *KibanaSpace) (*KibanaSpace, error) {

		if kibanaSpace == nil {
			return nil, NewAPIError(600, "You must provide kibana space object")
		}
		log.Debug("KibanaSpace: ", kibanaSpace)

		jsonData, err := json.Marshal(kibanaSpace)
		if err != nil {
			return nil, err
		}
		resp, err := c.R().SetBody(jsonData).Post(basePathKibanaSpace)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		kibanaSpace = &KibanaSpace{}
		err = json.Unmarshal(resp.Body(), kibanaSpace)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaSpace: ", kibanaSpace)

		return kibanaSpace, nil
	}

}

// newKibanaSpaceDeleteFunc permit to delete the kubana space wiht it id
func newKibanaSpaceDeleteFunc(c *resty.Client) KibanaSpaceDelete {
	return func(id string) error {

		if id == "" {
			return NewAPIError(600, "You must provide kibana space ID")
		}

		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s", basePathKibanaSpace, id)
		resp, err := c.R().Delete(path)
		if err != nil {
			return err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {

			return NewAPIError(resp.StatusCode(), resp.Status())

		}

		return nil
	}

}

// newKibanaSpaceUpdateFunc permit to update the Kibana space
func newKibanaSpaceUpdateFunc(c *resty.Client) KibanaSpaceUpdate {
	return func(kibanaSpace *KibanaSpace) (*KibanaSpace, error) {

		if kibanaSpace == nil {
			return nil, NewAPIError(600, "You must provide kibana space object")
		}
		log.Debug("KibanaSpace: ", kibanaSpace)

		jsonData, err := json.Marshal(kibanaSpace)
		if err != nil {
			return nil, err
		}
		path := fmt.Sprintf("%s/%s", basePathKibanaSpace, kibanaSpace.ID)
		resp, err := c.R().SetBody(jsonData).Put(path)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		kibanaSpace = &KibanaSpace{}
		err = json.Unmarshal(resp.Body(), kibanaSpace)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaSpace: ", kibanaSpace)

		return kibanaSpace, nil
	}

}
