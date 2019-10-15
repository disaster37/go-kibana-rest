package kibana

import (
	"encoding/json"

	"github.com/disaster37/go-kibana-rest/kbapi"
	"github.com/stretchr/testify/assert"
)

func (s *KBTestSuite) TestKibanaSaveObject() {

	// Create new index pattern
	dataJson := `{"attributes": {"title": "test-pattern-*"}}`
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(dataJson), &data)
	if err != nil {
		panic(err)
	}
	resp, err := s.client.API.KibanaSavedObject.Create(data, "index-pattern", "test", true, "default")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "test", resp["id"])
	assert.Equal(s.T(), "test-pattern-*", resp["attributes"].(map[string]interface{})["title"])

	// Create new index pattern in space
	dataJson = `{"attributes": {"title": "test-pattern-*"}}`
	data = make(map[string]interface{})
	err = json.Unmarshal([]byte(dataJson), &data)
	if err != nil {
		panic(err)
	}
	resp, err = s.client.API.KibanaSavedObject.Create(data, "index-pattern", "test", true, "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "test", resp["id"])
	assert.Equal(s.T(), "test-pattern-*", resp["attributes"].(map[string]interface{})["title"])

	// Get index pattern
	resp, err = s.client.API.KibanaSavedObject.Get("index-pattern", "test", "default")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "test", resp["id"])

	// Get index pattern from space
	resp, err = s.client.API.KibanaSavedObject.Get("index-pattern", "test", "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "test", resp["id"])

	// Search index pattern
	parameters := &kbapi.OptionalFindParameters{
		Search:       "test",
		SearchFields: []string{"title"},
		Fields:       []string{"id"},
	}
	resp, err = s.client.API.KibanaSavedObject.Find("index-pattern", "default", parameters)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	dataRes := resp["saved_objects"].([]interface{})[0].(map[string]interface{})
	assert.Equal(s.T(), "test", dataRes["id"].(string))

	// Search index pattern from space
	parameters = &kbapi.OptionalFindParameters{
		Search:       "test",
		SearchFields: []string{"title"},
		Fields:       []string{"id"},
	}
	resp, err = s.client.API.KibanaSavedObject.Find("index-pattern", "testacc", parameters)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	dataRes = resp["saved_objects"].([]interface{})[0].(map[string]interface{})
	assert.Equal(s.T(), "test", dataRes["id"].(string))

	// Update index pattern
	dataJson = `{"attributes": {"title": "test-pattern2-*"}}`
	err = json.Unmarshal([]byte(dataJson), &data)
	if err != nil {
		panic(err)
	}
	resp, err = s.client.KibanaSavedObject.Update(data, "index-pattern", "test", "default")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "test", resp["id"])
	assert.Equal(s.T(), "test-pattern2-*", resp["attributes"].(map[string]interface{})["title"])

	// Update index pattern in space
	dataJson = `{"attributes": {"title": "test-pattern2-*"}}`
	err = json.Unmarshal([]byte(dataJson), &data)
	if err != nil {
		panic(err)
	}
	resp, err = s.client.KibanaSavedObject.Update(data, "index-pattern", "test", "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "test", resp["id"])
	assert.Equal(s.T(), "test-pattern2-*", resp["attributes"].(map[string]interface{})["title"])

	// Export index pattern
	request := []map[string]string{
		map[string]string{
			"type": "index-pattern",
			"id":   "test",
		},
	}
	resp, err = s.client.KibanaSavedObject.Export(nil, request, true, "default")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)

	// Export index pattern from space
	request = []map[string]string{
		map[string]string{
			"type": "index-pattern",
			"id":   "test",
		},
	}
	resp, err = s.client.KibanaSavedObject.Export(nil, request, true, "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)

	// import index pattern
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	resp2, err := s.client.KibanaSavedObject.Import(b, true, "default")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp2)
	assert.Equal(s.T(), true, resp2["success"])

	// import index pattern in space
	b, err = json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	resp, err = s.client.KibanaSavedObject.Import(b, true, "testacc")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), true, resp["success"])

	// Delete index pattern
	err = s.client.API.KibanaSavedObject.Delete("index-pattern", "test", "default")
	assert.NoError(s.T(), err)
	resp, err = s.client.API.KibanaSavedObject.Get("index-pattern", "test", "default")
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), resp)

	// Delete index pattern in space
	err = s.client.API.KibanaSavedObject.Delete("index-pattern", "test", "testacc")
	assert.NoError(s.T(), err)
	resp, err = s.client.API.KibanaSavedObject.Get("index-pattern", "test", "testacc")
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), resp)
}
