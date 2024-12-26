package bnet

import (
	"context"
	"github.com/FuzzyStatic/blizzard/v3"
)

// ToDo rename
type RegionalClient struct {
	context  context.Context
	usClient blizzard.Client
	euClient blizzard.Client
	twClient blizzard.Client
}
