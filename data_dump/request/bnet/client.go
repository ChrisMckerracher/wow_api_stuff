package bnet

import (
	"context"
	"fmt"
	"github.com/FuzzyStatic/blizzard/v3"
	"net/http"
)

type RegionalClient struct {
	Context  context.Context
	usClient *blizzard.Client
	euClient *blizzard.Client
	twClient *blizzard.Client
}

type RegionalConfig struct {
	usConfig blizzard.Config
	euConfig blizzard.Config
	twConfig blizzard.Config
}

func NewRegionalClient(ctx context.Context, clientId string, clientSecret string) RegionalClient {
	configs := generateConfigs(clientId, clientSecret)

	// ToDo: error handling
	usClient, _ := blizzard.NewClient(configs.usConfig)
	euClient, _ := blizzard.NewClient(configs.euConfig)
	twClient, _ := blizzard.NewClient(configs.twConfig)

	usClient.AccessTokenRequest(ctx)
	// ToDo: panic next two lines???????????????????????????
	//euClient.AccessTokenRequest(ctx)
	//twClient.AccessTokenRequest(ctx)

	return RegionalClient{
		Context:  ctx,
		usClient: usClient,
		euClient: euClient,
		twClient: twClient,
	}

}

func generateConfigs(clientId string, clientSecret string) RegionalConfig {
	regionalConfig := RegionalConfig{}

	defaultBlizzardConfig := blizzard.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		HTTPClient:   http.DefaultClient,
		Locale:       blizzard.EnUS,
	}

	regionalConfig.usConfig = defaultBlizzardConfig
	regionalConfig.usConfig.Region = blizzard.US

	regionalConfig.euConfig = defaultBlizzardConfig
	regionalConfig.euConfig.Region = blizzard.EU

	regionalConfig.twConfig = defaultBlizzardConfig
	regionalConfig.twConfig.Region = blizzard.TW

	return regionalConfig

}

func (r *RegionalClient) Select(region blizzard.Region) (*blizzard.Client, error) {
	switch region {
	case blizzard.US:
		return r.usClient, nil
	case blizzard.EU:
		return r.euClient, nil
	case blizzard.TW:
		return r.twClient, nil
	default:
		return nil, fmt.Errorf("Invalid Region")
	}
}
