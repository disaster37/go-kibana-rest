package kbapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

// DataView is the DataView API object
type DataView struct {
	ID           string                 `json:"id,omitempty"`
	Name         string                 `json:"name"`
	Title        string                 `json:"title"`
	Namespaces   []string               `json:"namespaces,omitempty"`
	AllowNoIndex bool                   `json:"allowNoIndex,omitempty"`
	TypeMeta     map[string]interface{} `json:"typeMeta,omitempty"`
}

type dataViewWrapper struct {
	DataView DataView `json:"data_view"`
}

// KibanaDataViewGet permit to get data view
type KibanaDataViewGet func(id string, spaceID string) (*DataView, error)

// KibanaDataViewList permit to get all data views
type KibanaDataViewList func(spaceID string) ([]DataView, error)

// KibanaDataViewCreate permit to create data view
type KibanaDataViewCreate func(kibanaDataView *DataView, spaceID string) (*DataView, error)

// KibanaDataViewDelete permit to delete data view
type KibanaDataViewDelete func(id string, spaceID string) error

// KibanaDataViewUpdate permit to update data view
type KibanaDataViewUpdate func(kibanaDataView *DataView, spaceID string) (*DataView, error)

// String permit to return KibanaDataView object as JSON string
func (k *DataView) String() string {
	j, _ := json.Marshal(k)
	return string(j)
}

// newKibanaDataViewGetFunc permit to get the kibana data view with its id
func newKibanaDataViewGetFunc(c *resty.Client) KibanaDataViewGet {
	return func(id string, spaceID string) (*DataView, error) {

		if id == "" {
			return nil, NewAPIError(600, "You must provide kibana data view ID")
		}
		log.Debug("ID: ", id)

		path := generatePath("data_view", id, spaceID)
		resp, err := c.R().Get(path)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			if resp.StatusCode() == 404 {
				return nil, nil
			}
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}

		dv := &dataViewWrapper{DataView: DataView{}}
		err = json.Unmarshal(resp.Body(), dv)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaDataView: ", dv)

		return &dv.DataView, nil
	}

}

// newKibanaDataViewListFunc permit to get all Kibana data view
func newKibanaDataViewListFunc(c *resty.Client) KibanaDataViewList {
	return func(spaceID string) ([]DataView, error) {

		type dataViews struct {
			DataViews []DataView `json:"data_view"`
		}

		path := generatePath("", "", spaceID)
		resp, err := c.R().Get(path)
		if err != nil {
			return []DataView{}, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return []DataView{}, NewAPIError(resp.StatusCode(), resp.Status())
		}
		dataViewsList := dataViews{}
		err = json.Unmarshal(resp.Body(), &dataViewsList)
		if err != nil {
			return []DataView{}, err
		}
		log.Debug("DataViews: ", dataViewsList)

		return dataViewsList.DataViews, nil
	}

}

// newKibanaDataViewCreateFunc permit to create new Kibana data view
func newKibanaDataViewCreateFunc(c *resty.Client) KibanaDataViewCreate {
	return func(kibanaDataView *DataView, spaceID string) (*DataView, error) {

		if kibanaDataView == nil {
			return nil, NewAPIError(600, "You must provide kibana data view object")
		}
		log.Debug("KibanaDataView: ", kibanaDataView)

		dv := &dataViewWrapper{DataView: *kibanaDataView}

		jsonData, err := json.Marshal(dv)
		if err != nil {
			return nil, err
		}
		path := generatePath("data_view", "", spaceID)
		resp, err := c.R().SetBody(jsonData).Post(path)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", string(jsonData))
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		dv = &dataViewWrapper{}
		err = json.Unmarshal(resp.Body(), dv)
		if err != nil {
			return nil, err
		}

		log.Debug("KibanaDataView: ", dv)

		return &dv.DataView, nil
	}

}

// newKibanaDataViewDeleteFunc permit to delete the kibana data view with its id
func newKibanaDataViewDeleteFunc(c *resty.Client) KibanaDataViewDelete {
	return func(id string, spaceID string) error {

		if id == "" {
			return NewAPIError(600, "You must provide kibana data view ID")
		}

		log.Debug("ID: ", id)

		path := generatePath("data_view", id, spaceID)
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

// newKibanaDataViewUpdateFunc permit to update the Kibana data view
func newKibanaDataViewUpdateFunc(c *resty.Client) KibanaDataViewUpdate {
	return func(kibanaDataView *DataView, spaceID string) (*DataView, error) {

		if kibanaDataView == nil {
			return nil, NewAPIError(600, "You must provide kibana data view object")
		}
		log.Debug("KibanaDataView: ", kibanaDataView)

		path := generatePath("data_view", kibanaDataView.ID, spaceID)

		dv := &dataViewWrapper{
			DataView: *kibanaDataView,
		}
		dv.DataView.ID = ""

		jsonData, err := json.Marshal(dv)
		if err != nil {
			return nil, err
		}
		resp, err := c.R().SetBody(jsonData).Post(path)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		dv = &dataViewWrapper{}
		err = json.Unmarshal(resp.Body(), dv)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaDataView: ", dv)

		return &dv.DataView, nil
	}
}

// Generate the path for a data view method.
// resourceType is the resource type being acted upon, `data_view` for CRUD, empty for the List operation
// id is optional since the Create/List methods do not specify an ID
// spaceId is optional since an empty value means it's the default space
func generatePath(resourceType string, id string, spaceId string) string {
	// Set the correct base path based on whether a space is provided
	base := "/api/data_views"
	if spaceId != "" {
		base = fmt.Sprintf("/s/%s/api/data_views", spaceId)
	}
	// Prepend a slash to a non-empty ID
	if id != "" {
		id = fmt.Sprintf("/%s", id)
	}
	// Prepend a slash to a non-empty resourceType
	if resourceType != "" {
		resourceType = fmt.Sprintf("/%s", resourceType)
	}
	return fmt.Sprintf("%s%s%s", base, resourceType, id)
}
