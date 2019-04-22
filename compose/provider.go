package compose

import (
	composeapi "github.com/compose/gocomposeapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Compose Authentication Token",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"compose_deployment": resourceDeployment(),
		},

		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return composeapi.NewClient(d.Get("api_token").(string))
		},
	}
}
