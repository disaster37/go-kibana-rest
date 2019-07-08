package kibana

import (
	"github.com/disaster37/go-kibana-rest/kbapi"
	"github.com/stretchr/testify/assert"
)

func (s *KBTestSuite) TestKibanaRoleManagement() {

	// Create new role
	kibanaRole := &kbapi.KibanaRole{
		Name: "test",
		Elasticsearch: &kbapi.KibanaRoleElasticsearch{
			Indices: []kbapi.KibanaRoleElasticsearchIndice{
				kbapi.KibanaRoleElasticsearchIndice{
					Names: []string{
						"*",
					},
					Privileges: []string{
						"read",
					},
				},
			},
		},
		Kibana: []kbapi.KibanaRoleKibana{
			kbapi.KibanaRoleKibana{
				Base: []string{
					"read",
				},
			},
		},
	}
	kibanaRole, err := s.client.API.KibanaRoleManagement.CreateOrUpdate(kibanaRole)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), kibanaRole)
	assert.Equal(s.T(), "test", kibanaRole.Name)

	// Get role
	kibanaRole, err = s.client.API.KibanaRoleManagement.Get("test")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), kibanaRole)
	assert.Equal(s.T(), "test", kibanaRole.Name)

	// List role
	kibanaRoles, err := s.client.API.KibanaRoleManagement.List()
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), kibanaRoles)

	// Delete role
	err = s.client.API.KibanaRoleManagement.Delete("test")
	assert.NoError(s.T(), err)
	kibanaRole, err = s.client.API.KibanaRoleManagement.Get("test")
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), kibanaRole)

}
