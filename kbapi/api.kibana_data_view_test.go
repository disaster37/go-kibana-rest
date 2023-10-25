package kbapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (s *KBAPITestSuite) TestKibanaDataView() {

	cases := []struct {
		description  string
		resourceType string
		id           string
		spaceId      string
		expected     string
	}{
		{
			description:  "List in default space",
			resourceType: "",
			id:           "",
			spaceId:      "",
			expected:     "/api/data_views",
		},
		{
			description:  "List in named space",
			resourceType: "",
			id:           "",
			spaceId:      "space1",
			expected:     "/s/space1/api/data_views",
		},
		{
			description:  "Create in default space",
			resourceType: "data_view",
			id:           "",
			spaceId:      "",
			expected:     "/api/data_views/data_view",
		},
		{
			description:  "Create in named space",
			resourceType: "data_view",
			id:           "",
			spaceId:      "space1",
			expected:     "/s/space1/api/data_views/data_view",
		},
		{
			description:  "Get/Update/Delete in default space",
			resourceType: "data_view",
			id:           "id123",
			spaceId:      "",
			expected:     "/api/data_views/data_view/id123",
		},
		{
			description:  "Get/Update/Delete in named space",
			resourceType: "data_view",
			id:           "id123",
			spaceId:      "space1",
			expected:     "/s/space1/api/data_views/data_view/id123",
		},
	}
	for _, tt := range cases {
		s.T().Run(tt.description, func(t *testing.T) {
			result := generatePath(tt.resourceType, tt.id, tt.spaceId)
			if result != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, result)
			}
		})
	}

	// Create new data view
	data := &DataView{
		ID:    "test-id",
		Name:  "test-name",
		Title: "test-title-*",
	}
	resp, err := s.API.KibanaDataViews.Create(data, "")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), data.ID, resp.ID)
	assert.Equal(s.T(), data.Name, resp.Name)
	assert.Equal(s.T(), data.Title, resp.Title)

	// Create new data view in a space
	spacedData := &DataView{
		ID:    "testacc-test-id",
		Name:  "testacc-test-name",
		Title: "testacc-test-title-*",
	}
	resp, err = s.API.KibanaDataViews.Create(spacedData, "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), spacedData.ID, resp.ID)
	assert.Equal(s.T(), spacedData.Name, resp.Name)
	assert.Equal(s.T(), spacedData.Title, resp.Title)

	// List all data views
	listResp, err := s.API.KibanaDataViews.List("")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), listResp)
	found := 0
	for _, view := range listResp {
		if view.ID == data.ID {
			assert.Equal(s.T(), data.Name, view.Name)
			assert.Equal(s.T(), data.Title, view.Title)
			found++
		}
	}
	assert.Equal(s.T(), 1, found)

	// List all data views in a space
	listResp, err = s.API.KibanaDataViews.List("testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), listResp)
	found = 0
	for _, view := range listResp {
		if view.ID == spacedData.ID {
			assert.Equal(s.T(), spacedData.Name, view.Name)
			assert.Equal(s.T(), spacedData.Title, view.Title)
			found++
		}
	}
	assert.Equal(s.T(), 1, found)

	// Get specific data view
	resp, err = s.API.KibanaDataViews.Get(data.ID, "")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), data.ID, resp.ID)

	// Get specific data view in a space
	resp, err = s.API.KibanaDataViews.Get(spacedData.ID, "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), spacedData.ID, resp.ID)

	// Update data view
	updatedData := &DataView{
		ID:    data.ID,
		Name:  "test-name-updated",
		Title: "test-title-updated-*",
	}
	resp, err = s.API.KibanaDataViews.Update(updatedData, "")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), updatedData.ID, resp.ID)
	assert.Equal(s.T(), updatedData.Name, resp.Name)
	assert.Equal(s.T(), updatedData.Title, resp.Title)

	// Update data view in a space
	spacedUpdatedData := &DataView{
		ID:    spacedData.ID,
		Name:  "testacc-test-name-updated",
		Title: "testacc-test-title-updated-*",
	}
	resp, err = s.API.KibanaDataViews.Update(spacedUpdatedData, "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), spacedUpdatedData.ID, resp.ID)
	assert.Equal(s.T(), spacedUpdatedData.Name, resp.Name)
	assert.Equal(s.T(), spacedUpdatedData.Title, resp.Title)

	// Delete data view
	err = s.API.KibanaDataViews.Delete(updatedData.ID, "")
	assert.NoError(s.T(), err)
	resp, err = s.API.KibanaDataViews.Get(updatedData.ID, "")
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), resp)

	// Delete data view in a space
	err = s.API.KibanaDataViews.Delete(spacedUpdatedData.ID, "testacc")
	assert.NoError(s.T(), err)
	resp, err = s.API.KibanaDataViews.Get(spacedUpdatedData.ID, "testacc")
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), resp)
}
