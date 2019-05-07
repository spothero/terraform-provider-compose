// Copyright 2019 SpotHero
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
