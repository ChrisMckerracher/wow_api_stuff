package character

import (
	"context"
	"data_dump/request/bnet"
	"git
	"github.com/stretchr/testify/suite"
	"context"
	"testing"
)

type CharacterTestSuite struct {
	suite.Suite
	sut Client
}

func (c *CharacterTestSuite) SetupSuite() {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	regionalClient := bnet.NewRegionalClient(context.TODO(), clientId, clientSecret)
	c.sut = Client{
		RegionalClient: regionalClient,
	}

}

func (c *CharacterTestSuite) TestBasiclyNothing() {
	c.sut.Scrape()
	//
}

func TestCharacters(t *testing.T) {
	suite.Run(t, new(CharacterTestSuite))
}