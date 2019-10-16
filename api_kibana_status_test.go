package kibana

import (
	"github.com/stretchr/testify/assert"
)

func (s *KBTestSuite) TestKibanaStatus() {

	// List kibana space
	kibanaStatus, err := s.client.API.KibanaStatus.Get()
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), kibanaStatus)
}
