package provider

import (
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
				ForceNew: true,
			},
		},
	}
}

func configureVictoropsTeam(d *schema.ResourceData, t models.TeamDetail) {
	d.SetId(t.Slug)
	d.Set("name", t.Name)
	d.Set("slug", t.Slug)
	d.Set("version", t.Version)
	d.Set("self_url", t.SelfURL)
}

func resourceTeamCreate(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.PostAPIPublicV1TeamParams{
		Context: meta.StopContext,
		Body: &models.AddTeamPayload{
			Name: strPointer(d.Get("name").(string)),
		},
	}
	res, err := meta.Client.Operations.PostAPIPublicV1Team(params)
	if err != nil {
		return err
	} else if res == nil || res.Payload == nil {
		return errors.New("API Error")
	}
	configureVictoropsTeam(d, res.Payload.TeamDetail)
	return nil
}

func resourceTeamRead(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.GetAPIPublicV1TeamTeamParams{
		Context: meta.StopContext,
		Team:    strPointer(d.Id()),
	}
	res, err := meta.Client.Operations.GetAPIPublicV1TeamTeam(params)
	if err != nil {
		return err
	} else if res == nil || res.Payload == nil {
		return errors.New("API Error")
	}
	configureVictoropsTeam(d, res.Payload.TeamDetail)
	return nil
}

func resourceTeamUpdate(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not supported")
}

func resourceTeamDelete(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.DeleteAPIPublicV1TeamTeamParams{
		Context: meta.StopContext,
		Team:    strPointer(d.Id()),
	}
	res, err := meta.Client.Operations.DeleteAPIPublicV1TeamTeam(params)
	if err != nil {
		return err
	} else if res == nil {
		return errors.New("API Error")
	}
	return nil
}
