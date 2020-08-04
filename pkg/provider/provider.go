package provider

import (
	"context"
	"terraform-provider-victorops/pkg/api/client"
	"terraform-provider-victorops/pkg/api/client/operations"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

type Meta struct {
	Client      *client.Victorops
	StopContext context.Context
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.victorops.com/",
			},
			"api_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"victorops_team": resourceTeam(),
			"victorops_user": resourceUser(),
		},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			URL: d.Get("api_url").(string),
			ID:  d.Get("api_id").(string),
			Key: d.Get("api_key").(string),
		}

		client, err := config.NewClient()
		if err != nil {
			return nil, err
		}

		params := &operations.GetAPIPublicV1MaintenancemodeParams{
			Context: p.StopContext(),
		}
		res, err := client.Operations.GetAPIPublicV1Maintenancemode(params)
		if err != nil {
			return nil, errors.Wrapf(err, "Health check API call to VictorOps API failed")
		} else if res == nil || res.Payload == nil {
			return nil, errors.New("Health check API call to VictorOps API failed")
		}

		return &Meta{
			Client:      client,
			StopContext: p.StopContext(),
		}, nil
	}
}
