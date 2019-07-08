package kibana

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func (s *KBTestSuite) TestKibanaDashboard() {

	// Import dashboard from fixtures
	b, err := ioutil.ReadFile("fixtures/kibana-dashboard.json")
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(b, &data)
	err = s.client.API.KibanaDashboard.Import(data, nil, true)
	assert.NoError(s.T(), err)

	// Export dashboard
	data, err = s.client.API.KibanaDashboard.Export([]string{"edf84fe0-e1a0-11e7-b6d5-4dc382ef7f5b"})
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), data)

}
