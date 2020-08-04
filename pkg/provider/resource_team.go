package provider

import (
	"fmt"
	"terraform-provider-victorops/pkg/api/client/operations"
	"terraform-provider-victorops/pkg/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeamCreate,
		Read:   resourceTeamRead,
		Update: resourceTeamUpdate,
		Delete: resourceTeamDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTeamConfigure(d *schema.ResourceData, t models.TeamDetail) {
	d.SetId(t.Slug)
	d.Set("name", t.Name)
	d.Set("slug", t.Slug)
	d.Set("version", t.Version)
	d.Set("self_url", t.SelfURL)
}

func resourceTeamToString(d *schema.ResourceData) string {
	return fmt.Sprintf("victorops_team: %s", d.Get("name"))
}

func resourceTeamCreate(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.PostAPIPublicV1TeamParams{
		Context: meta.StopContext,
		Body: &models.AddTeamPayload{
			Name: &[]string{d.Get("name").(string)}[0],
		},
	}
	res, err := meta.Client.Operations.PostAPIPublicV1Team(params)
	if err != nil {
		return errors.Wrapf(err, "Error creating resource %s: %v", resourceTeamToString(d), res)
	} else if res == nil || res.Payload == nil {
		return errors.Errorf("Unknown API response creating resource %s", resourceTeamToString(d))
	}
	resourceTeamConfigure(d, res.Payload.TeamDetail)
	return nil
}

func resourceTeamRead(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.GetAPIPublicV1TeamTeamParams{
		Context: meta.StopContext,
		Team:    &[]string{d.Id()}[0],
	}
	res, err := meta.Client.Operations.GetAPIPublicV1TeamTeam(params)
	if err != nil {
		return errors.Wrapf(err, "Error updating resource %s: %v", resourceTeamToString(d), res)
	} else if res == nil || res.Payload == nil {
		return errors.Errorf("Unknown API response reading resource %s", resourceTeamToString(d))
	}
	resourceTeamConfigure(d, res.Payload.TeamDetail)
	return nil
}

func resourceTeamUpdate(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not supported")
}

func resourceTeamDelete(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.DeleteAPIPublicV1TeamTeamParams{
		Context: meta.StopContext,
		Team:    &[]string{d.Id()}[0],
	}
	res, err := meta.Client.Operations.DeleteAPIPublicV1TeamTeam(params)
	if err != nil {
		return errors.Wrapf(err, "Error deleting resource %s: %v", resourceTeamToString(d), res)
	} else if res == nil {
		return errors.Errorf("Unknown API response deleting resource %s", resourceTeamToString(d))
	}
	return nil
}
