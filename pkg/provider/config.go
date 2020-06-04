package provider

import (
	"net/url"
	"terraform-provider-victorops/pkg/api/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Config struct {
	URL string
	ID  string
	Key string
}

type AuthWriter struct {
	config *Config
}

func (w *AuthWriter) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	req.SetHeaderParam("X-VO-Api-Id", w.config.ID)
	req.SetHeaderParam("X-VO-Api-Key", w.config.Key)
	return nil
}

func (c *Config) NewClient() (*client.Victorops, error) {
	uri, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return nil, err
	}
	transport := httptransport.New(uri.Host, uri.Path, []string{uri.Scheme})
	transport.DefaultAuthentication = &AuthWriter{config: c}
	client := client.New(transport, nil)
	return client, nil
}
