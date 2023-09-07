package kibana

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type KBTestSuite struct {
	suite.Suite
}

func (s *KBTestSuite) SetupSuite() {

	// Init logger
	logrus.SetFormatter(new(prefixed.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)

}

func (s *KBTestSuite) SetupTest() {

	// Do somethink before each test

}

func TestKBTestSuite(t *testing.T) {
	suite.Run(t, new(KBTestSuite))
}

func (s *KBTestSuite) TestNewClient() {

	cfg := Config{
		Address:          "http://127.0.0.1:5601",
		Username:         "elastic",
		Password:         "changeme",
		DisableVerifySSL: true,
	}

	client, err := NewClient(cfg)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), client)
	assert.NotNil(s.T(), "Basic", client.Client.AuthScheme)
	assert.Equal(s.T(), "elastic", client.Client.UserInfo.Username)
	assert.Equal(s.T(), "changeme", client.Client.UserInfo.Password)
}

func (s *KBTestSuite) TestNewAPIKeyClient() {

	cfg := Config{
		Address:          "http://127.0.0.1:5601",
		ApiKey:           "foo",
		DisableVerifySSL: true,
	}

	client, err := NewClient(cfg)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), client)

	assert.Equal(s.T(), "ApiKey", client.Client.AuthScheme)
	assert.Equal(s.T(), "foo", client.Client.Token, "foo")

}

func (s *KBTestSuite) TestNewDefaultClient() {

	client, err := NewDefaultClient()

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), client)

	assert.NotNil(s.T(), client.Client.AuthScheme, "Basic")
	assert.Equal(s.T(), client.Client.UserInfo.Username, "")
	assert.Equal(s.T(), client.Client.UserInfo.Password, "")
	assert.Equal(s.T(), client.Client.Token, "")
}
