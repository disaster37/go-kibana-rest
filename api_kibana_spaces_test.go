package kibana

import (
	"github.com/disaster37/go-kibana-rest/kbapi"
	"github.com/stretchr/testify/assert"
)

func (s *KBTestSuite) TestKibanaSpaces() {

	// List kibana space
	kibanaSpaces, err := s.client.API.KibanaSpaces.List()
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), kibanaSpaces)

	// Get the default Space
	kibanaSpace, err := s.client.API.KibanaSpaces.Get(kibanaSpaces[0].ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), kibanaSpaces[0].ID, kibanaSpace.ID)
	assert.Equal(s.T(), "Default", kibanaSpace.Name)

	// Create new space
	kibanaSpace = &kbapi.KibanaSpace{
		ID:          "test",
		Name:        "test",
		Description: "My test",
	}
	kibanaSpace, err = s.client.KibanaSpaces.Create(kibanaSpace)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), kibanaSpace.ID)

	// Update space
	kibanaSpace.Name = "test2"
	kibanaSpace, err = s.client.KibanaSpaces.Update(kibanaSpace)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "test2", kibanaSpace.Name)

	// Copy object on space
	parameter := &kbapi.KibanaSpaceCopySavedObjectParameter{
		Spaces:            []string{"test"},
		IncludeReferences: true,
		Overwrite:         true,
		Objects: []kbapi.KibanaSpaceObjectParameter{
			kbapi.KibanaSpaceObjectParameter{
				Type: "config",
				Id:   "7.4.0",
			},
		},
	}
	err = s.client.KibanaSpaces.CopySavedObjects(parameter, "")
	assert.NoError(s.T(), err)

	// Delete space
	err = s.client.KibanaSpaces.Delete(kibanaSpace.ID)
	assert.NoError(s.T(), err)
	kibanaSpace, err = s.client.KibanaSpaces.Get(kibanaSpace.ID)
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), kibanaSpace)

}
