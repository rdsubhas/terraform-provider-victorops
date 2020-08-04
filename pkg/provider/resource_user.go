package provider

import (
	"fmt"
	"terraform-provider-victorops/pkg/api/client/operations"
	"terraform-provider-victorops/pkg/api/models"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceUserConfigure(d *schema.ResourceData, u models.V1User) {
	d.SetId(u.Username)
	d.Set("username", u.Username)
	d.Set("first_name", u.FirstName)
	d.Set("last_name", u.LastName)
	d.Set("email", u.Email)
}

func resourceUserToString(d *schema.ResourceData) string {
	return fmt.Sprintf("victorops_user: %s (%s)", d.Get("username"), d.Get("email"))
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.PostAPIPublicV1UserParams{
		Context: meta.StopContext,
		Body: &models.AddUserPayload{
			FirstName: &[]string{d.Get("first_name").(string)}[0],
			LastName:  &[]string{d.Get("last_name").(string)}[0],
			Username:  &[]string{d.Get("username").(string)}[0],
			Email:     &[]strfmt.Email{strfmt.Email(d.Get("email").(string))}[0],
		},
	}
	res, err := meta.Client.Operations.PostAPIPublicV1User(params)
	if err != nil {
		return errors.Wrapf(err, "Error creating resource %s: %v", resourceUserToString(d), res)
	} else if res == nil || res.Payload == nil {
		return errors.Errorf("Unknown API response creating resource %s", resourceUserToString(d))
	}
	resourceUserConfigure(d, res.Payload.V1User)
	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.GetAPIPublicV1UserUserParams{
		Context: meta.StopContext,
		User:    &[]string{d.Id()}[0],
	}
	res, err := meta.Client.Operations.GetAPIPublicV1UserUser(params)
	if err != nil {
		return errors.Wrapf(err, "Error updating resource %s: %v", resourceUserToString(d), res)
	} else if res == nil || res.Payload == nil {
		return errors.Errorf("Unknown API response reading resource %s", resourceUserToString(d))
	}
	resourceUserConfigure(d, res.Payload.V1User)
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return errors.New("Not supported")
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	meta := m.(*Meta)
	params := &operations.DeleteAPIPublicV1UserUserParams{
		Context: meta.StopContext,
		User:    &[]string{d.Id()}[0],
		Body: &models.DeleteUserPayload{
			ReplacementStrategy: "teamAdmin",
		},
	}
	res, err := meta.Client.Operations.DeleteAPIPublicV1UserUser(params)
	if err != nil {
		return errors.Wrapf(err, "Error deleting resource %s: %v", resourceUserToString(d), res)
	} else if res == nil {
		return errors.Errorf("Unknown API response deleting resource %s", resourceUserToString(d))
	}
	return nil
}
