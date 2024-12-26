package bnet

import (
	"context"
	"fmt"
	"github.com/FuzzyStatic/blizzard/v3"
)

type RegionalClient struct {
	Context  context.Context
	usClient blizzard.Client
	euClient blizzard.Client
	twClient blizzard.Client
}

type Region int

const (
	US Region = iota
	EU
	TW
)

func (r *RegionalClient) Select(region Region) (*blizzard.Client, error) {
	switch region {
	case US:
		return &r.usClient, nil
	case EU:
		return &r.euClient, nil
	case TW:
		return &r.twClient, nil
	default:
		return nil, fmt.Errorf("Invalid Region")
	}
}
