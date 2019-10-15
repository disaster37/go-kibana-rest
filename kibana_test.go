package kibana

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/disaster37/go-kibana-rest/kbapi"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type KBTestSuite struct {
	suite.Suite
	client *Client
}

func (s *KBTestSuite) SetupSuite() {

	// Init logger
	logrus.SetFormatter(new(prefixed.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	address := os.Getenv("KIBANA_URL")
	username := os.Getenv("KIBANA_USERNAME")
	password := os.Getenv("KIBANA_PASSWORD")

	if address == "" {
		panic("You need to put kibana url on environment variable KIBANA_URL. If you need auth, you can use KIBANA_USERNAME and KIBANA_PASSWORD")
	}

	// Init client
	config := Config{
		Address:  address,
		Username: username,
		Password: password,
	}

	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	// Wait kb is online
	isOnline := false
	nbTry := 0
	for isOnline == false {
		_, err := client.API.KibanaSpaces.List()
		if err == nil {
			isOnline = true
		} else {
			time.Sleep(5 * time.Second)
			if nbTry == 10 {
				panic(fmt.Sprintf("We wait 50s that Kibana start: %s", err))
			}
			nbTry++
		}
	}

	// Create kibana space
	space := &kbapi.KibanaSpace{
		ID:   "testacc",
		Name: "testacc",
	}
	_, err = client.API.KibanaSpaces.Create(space)
	if err != nil {
		if err.(kbapi.APIError).Code != 409 {
			panic(err)
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
