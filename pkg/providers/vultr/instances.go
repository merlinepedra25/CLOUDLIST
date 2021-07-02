package vultr

import (
	"context"

	"github.com/projectdiscovery/cloudlist/pkg/schema"
	"github.com/vultr/govultr/v2"
)

// instanceProvider is an instance provider for vultr API
type instanceProvider struct {
	profile string
	client  *govultr.Client
}

// GetInstances returns all the instances in the store for a provider.
func (d *instanceProvider) GetResource(ctx context.Context) (*schema.Resources, error) {
	list := &schema.Resources{}

	listOptions := &govultr.ListOptions{PerPage: 1}
	for {
		intances, meta, err := d.client.Instance.List(ctx, listOptions)
		if err != nil {
			return nil, err
		}
		for _, inst := range intances {
			list.Append(&schema.Resource{
				Provider:   providerName,
				PublicIPv4: inst.MainIP,
				Profile:    d.profile,
				Public:     inst.MainIP != "",
			})
		}

		if meta.Links.Next == "" {
			break
		} else {
			listOptions.Cursor = meta.Links.Next
			continue
		}
	}

	return list, nil
}
