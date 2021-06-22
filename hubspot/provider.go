package hubspot

import (
	"terraform-provider-hubspot/client"
	"terraform-provider-hubspot/token"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HUBSPOT_CLIENT_ID", nil),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HUBSPOT_CLIENT_SECRET", nil),
			},
			"refresh_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HUBSPOT_REFRESH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hubspot_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hubspot_user": dataSourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	accessToken := token.GenerateToken(d.Get("client_id").(string), d.Get("client_secret").(string), d.Get("refresh_token").(string))
	return client.NewClient(accessToken), nil
}
