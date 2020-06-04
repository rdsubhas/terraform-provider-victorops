package provider

import (
	"context"
	"terraform-provider-victorops/pkg/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Meta struct {
	Client      *client.Victorops
	StopContext context.Context
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			URL: d.Get("apiurl").(string),
			ID:  d.Get("apiId").(string),
			Key: d.Get("apiKey").(string),
		}

		client, err := config.NewClient()
		if err != nil {
			return nil, err
		}

		return &Meta{
			Client:      client,
			StopContext: p.StopContext(),
		}, nil
	}
}
