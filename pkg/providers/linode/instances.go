package linode

import (
	"context"

	"github.com/linode/linodego"
	"github.com/projectdiscovery/cloudlist/pkg/schema"
)

// instanceProvider is an instance provider for linode API
type instanceProvider struct {
	profile string
	client  *linodego.Client
}

// GetResource returns all the instance resources for a provider.
func (d *instanceProvider) GetResource(ctx context.Context) (*schema.Resources, error) {

	// Using autop-agination as mentioned by https://github.com/linode/linodego#auto-pagination-requests
	// We can also use handle pagination manually if needed
	instances, err := d.client.ListInstances(ctx, nil)
	if err != nil {
		return nil, err
	}

	list := &schema.Resources{}

	for _, inst := range instances {

		// Assuming (and obseved the same) first IP in the list is the public IP
		ip4 := inst.IPv4[0].String()

		list.Append(&schema.Resource{
			Provider:   providerName,
			PublicIPv4: ip4,
			Profile:    d.profile,
			Public:     ip4 != "",
		})
	}

	return list, nil
}
