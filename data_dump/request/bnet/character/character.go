package character

import (
	"data_dump/request/bnet"
	"fmt"
	
)

// todo: move these
type Character string

// ToDo: typify realms, character and realm types are cross api also :)
type Server string

type Client struct {
	bnet.RegionalClient
}

// Scrape scrapes the leaderboard
func (c *Client) Scrape() (interface{}, error) {
	bnetClient, _ := c.Select(blizzard.US)
	answer, _, err := bnetClient.WoWMythicKeystoneLeaderboard(c.Context, 11, 197, 641)
	if answer != nil {
		fmt.Println(answer.LeadingGroups)
	}
	return answer, err
}
