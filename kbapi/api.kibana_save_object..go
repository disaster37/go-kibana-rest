package kbapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaSavedObject = "/api/saved_objects" // Base URL to access on Kibana save objects
)

type KibanaSavedObjectGet func(objectType string, id string) (map[string]interface{}, error)
type KibanaSavedObjectCreate func(data map[string]interface{}, objectType string, id string, overwrite bool) error
type KibanaSavedObjectUpdate func(data map[string]interface{}, objectType string, id string) error
type KibanaSavedObjectDelete func(objectType string, id string) error

// newKibanaSavedObjectGetFunc permit to get saved obejct by it id and type
func newKibanaSavedObjectGetFunc(c *resty.Client) KibanaSavedObjectGet {
	return func(objectType string, id string) (map[string]interface{}, error) {

		if objectType == "" {
			return nil, NewAPIError(600, "You must provide the object type")
		}
		if id == "" {
			return nil, NewAPIError(600, "You must provide the object ID")
		}
		log.Debug("ObjectType: ", objectType)
		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s/%s", basePathKibanaSavedObject, objectType, id)
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
		var data map[string]interface{}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return nil, err
		}
		log.Debug("Data: ", data)

		return data, nil
	}

}

// newKibanaSavedObjectCreateFunc permit to create new object on Kibana
func newKibanaSavedObjectCreateFunc(c *resty.Client) KibanaSavedObjectCreate {
	return func(data map[string]interface{}, objectType string, id string, overwrite bool) error {

		if data == nil {
			return NewAPIError(600, "You must provide one or more dashboard to import")
		}
		if objectType == "" {
			return NewAPIError(600, "You must provide the object type")
		}
		log.Debug("data: ", data)
		log.Debug("ObjectType: ", objectType)
		log.Debug("ID: ", id)
		log.Debug("Overwrite: ", overwrite)

		path := fmt.Sprintf("%s/%s%s", basePathKibanaSavedObject, objectType, id)
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		resp, err := c.R().SetQueryString(fmt.Sprintf("overwrite=%t", overwrite)).SetBody(jsonData).Post(path)
		if err != nil {
			return err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return NewAPIError(resp.StatusCode(), resp.Status())
		}
		var dataResponse map[string]interface{}
		err = json.Unmarshal(resp.Body(), &dataResponse)
		if err != nil {
			return err
		}
		log.Debug("Data response: ", dataResponse)

		return nil
	}
}

// newKibanaSavedObjectUpdateFunc permit to update object on Kibana
func newKibanaSavedObjectUpdateFunc(c *resty.Client) KibanaSavedObjectUpdate {
	return func(data map[string]interface{}, objectType string, id string) error {

		if data == nil {
			return NewAPIError(600, "You must provide one or more dashboard to import")
		}
		if objectType == "" {
			return NewAPIError(600, "You must provide the object type")
		}
		if id == "" {
			return NewAPIError(600, "You must provide the ID")
		}
		log.Debug("data: ", data)
		log.Debug("ObjectType: ", objectType)
		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s/%s", basePathKibanaSavedObject, objectType, id)
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		resp, err := c.R().SetBody(jsonData).Put(path)
		if err != nil {
			return err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return NewAPIError(resp.StatusCode(), resp.Status())
		}
		var dataResponse map[string]interface{}
		err = json.Unmarshal(resp.Body(), &dataResponse)
		if err != nil {
			return err
		}
		log.Debug("Data response: ", dataResponse)

		return nil
	}
}

// newKibanaSavedObjectDeleteFunc permit to delete object on Kibana
func newKibanaSavedObjectDeleteFunc(c *resty.Client) KibanaSavedObjectDelete {
	return func(objectType string, id string) error {

		if objectType == "" {
			return NewAPIError(600, "You must provide the object type")
		}
		if id == "" {
			return NewAPIError(600, "You must provide the id")
		}
		log.Debug("objectType: ", objectType)
		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s/%s", basePathKibanaSavedObject, objectType, id)
		resp, err := c.R().Delete(path)
		if err != nil {
			return err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return NewAPIError(resp.StatusCode(), resp.Status())
		}
		var dataResponse map[string]interface{}
		err = json.Unmarshal(resp.Body(), &dataResponse)
		if err != nil {
			return err
		}
		log.Debug("Data response: ", dataResponse)

		return nil
	}
}
