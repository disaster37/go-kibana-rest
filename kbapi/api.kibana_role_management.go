package kbapi

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaRoleManagement = "/api/security/role" // Base URL to access on Kibana role management
)

// Kibana role management object
type KibanaRole struct {
	Name            string                       `json:"name,omitempty"`
	Metadata        map[string]interface{}       `json:"metadata,omitempty"`
	TransientMedata *KibanaRoleTransientMetadata `json:"transient_metadata,omitempty"`
	Elasticsearch   *KibanaRoleElasticsearch     `json:"elasticsearch,omitempty"`
	Kibana          []KibanaRoleKibana           `json:"kibana,omitempty"`
}

type KibanaRoleTransientMetadata struct {
	Enabled bool `json:"enabled,omitempty"`
}
type KibanaRoleElasticsearch struct {
	Indices []KibanaRoleElasticsearchIndice `json:"indices,omitempty"`
	Cluster []string                        `json:"cluster,omitempty"`
	RunAs   []string                        `json:"run_as,omitempty"`
}
type KibanaRoleKibana struct {
	Base    []string            `json:"base,omitempty"`
	Feature map[string][]string `json:"feature,omitempty"`
	Spaces  []string            `json:"spaces,omitempty"`
}
type KibanaRoleElasticsearchIndice struct {
	Names         []string            `json:"names,omitempty"`
	Privileges    []string            `json:"privileges,omitempty"`
	FieldSecurity map[string][]string `json:"field_security,omitempty"`
	Query         interface{}         `json:"query,omitempty"`
}

func (k *KibanaRole) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}

// List of KibanaRole objects
type KibanaRoles []KibanaRole

type KibanaRoleManagementGet func(name string) (*KibanaRole, error)
type KibanaRoleManagementList func() (KibanaRoles, error)
type KibanaRoleManagementCreateOrUpdate func(kibanaRole *KibanaRole) (*KibanaRole, error)
type KibanaRoleManagementDelete func(name string) error

// newKibanaRoleManagementGetFunc permit to get the kibana role with it name
func newKibanaRoleManagementGetFunc(c *resty.Client) KibanaRoleManagementGet {
	return func(name string) (*KibanaRole, error) {

		if name == "" {
			return nil, NewAPIError(600, "You must provide kibana role name")
		}
		log.Debug("Name: ", name)

		path := fmt.Sprintf("%s/%s", basePathKibanaRoleManagement, name)
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
		kibanaRole := &KibanaRole{}
		err = json.Unmarshal(resp.Body(), kibanaRole)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaRole: ", kibanaRole)

		return kibanaRole, nil
	}

}

// newKibanaRoleManagementListFunc permit to get all kibana role
func newKibanaRoleManagementListFunc(c *resty.Client) KibanaRoleManagementList {
	return func() (KibanaRoles, error) {

		resp, err := c.R().Get(basePathKibanaRoleManagement)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		kibanaRoles := make(KibanaRoles, 0, 1)
		err = json.Unmarshal(resp.Body(), &kibanaRoles)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaRoles: ", kibanaRoles)

		return kibanaRoles, nil
	}

}

// newKibanaRoleManagementGetFunc permit to create or update the kibana role
func newKibanaRoleManagementCreateOrUpdateFunc(c *resty.Client) KibanaRoleManagementCreateOrUpdate {
	return func(kibanaRole *KibanaRole) (*KibanaRole, error) {

		if kibanaRole == nil {
			return nil, NewAPIError(600, "You must provide kibana role object")
		}
		log.Debug("Kibana role: ", kibanaRole)
		roleName := kibanaRole.Name

		path := fmt.Sprintf("%s/%s", basePathKibanaRoleManagement, roleName)
		kibanaRole.Name = ""
		jsonData, err := json.Marshal(kibanaRole)
		log.Debugf("Payload: %s", jsonData)
		if err != nil {
			return nil, err
		}
		resp, err := c.R().SetBody(jsonData).Put(path)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}

		// Retrive the object to return it
		kibanaRole, err = newKibanaRoleManagementGetFunc(c)(roleName)
		if err != nil {
			return nil, err
		}

		log.Debug("KibanaRole: ", kibanaRole)

		return kibanaRole, nil
	}

}

// newKibanaRoleManagementDeleteFunc permit to delete kibana role with it name
func newKibanaRoleManagementDeleteFunc(c *resty.Client) KibanaRoleManagementDelete {
	return func(name string) error {

		if name == "" {
			return NewAPIError(600, "You must provide kibana role name")
		}
		log.Debug("Name: ", name)

		path := fmt.Sprintf("%s/%s", basePathKibanaRoleManagement, name)
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
