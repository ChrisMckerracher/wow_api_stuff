package character

import (
	"data_dump/request/bnet"
)

// todo: move these
type Character string

// ToDo: typify realms, character and realm types are cross api also :)
type Server string

type Client struct {
	bnet.RegionalClient
}

func (c *Client) Scrape() {

}

func (c *Client) ScrapeUser(user Character, realm Server, region bnet.Region) {
	bnetClient, err := c.Select(region)

	if err != nil {

	}

	//bnetClient.WoWCharacterMythicKeystoneProfile(c.Context, string(realm), string(user))
}
