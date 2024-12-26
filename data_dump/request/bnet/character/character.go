package character

import (
	"data_dump/request/bnet"
)

type Client struct {
	bnet.RegionalClient
}

func (c *Client) Scrape() {

}
