package vultr

import (
	"context"

	"github.com/projectdiscovery/cloudlist/pkg/schema"
	"github.com/vultr/govultr/v2"
	"golang.org/x/oauth2"
)

const (
	providerName = "vultr"
	apiKey       = "vultr_api_key"
)

// Provider is a data provider for vultr API
type Provider struct {
	profile string
	client  *govultr.Client
}

// New creates a new provider client for vultr API
func New(options schema.OptionBlock) (*Provider, error) {
	apiKey, ok := options.GetMetadata(apiKey)
	if !ok {
		return nil, &schema.ErrNoSuchKey{Name: apiKey}
	}
	profile, _ := options.GetMetadata("profile")

	config := &oauth2.Config{}

	ctx := context.Background()
	TokenSource := config.TokenSource(ctx, &oauth2.Token{AccessToken: apiKey})
	vultrClient := govultr.NewClient(oauth2.NewClient(ctx, TokenSource))

	return &Provider{profile: profile, client: vultrClient}, nil
}

// Name returns the name of the provider
func (p *Provider) Name() string {
	return providerName
}

// ProfileName returns the name of the provider profile
func (p *Provider) ProfileName() string {
	return p.profile
}

// Resources returns the provider for a deployment source.
func (p *Provider) Resources(ctx context.Context) (*schema.Resources, error) {
	provider := &instanceProvider{client: p.client, profile: p.profile}
	return provider.GetResource(ctx)
}
