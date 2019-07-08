package kibana

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"testing"
	"time"
)

type KBTestSuite struct {
	suite.Suite
	client *Client
}

func (s *KBTestSuite) SetupSuite() {

	// Init logger
	logrus.SetFormatter(new(prefixed.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	// Init client
	config := Config{
		Address:  "http://kb:5601",
		Username: "elastic",
		Password: "changeme",
	}

	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	// Wait kb is online
	isOnline := false
	for isOnline == false {
		_, err := client.API.KibanaSpaces.List()
		if err == nil {
			isOnline = true
		} else {
			time.Sleep(5 * time.Second)
		}
	}

	s.client = client

}

func (s *KBTestSuite) SetupTest() {

	// Do somethink before each test

}

func TestKBTestSuite(t *testing.T) {
	suite.Run(t, new(KBTestSuite))
}
