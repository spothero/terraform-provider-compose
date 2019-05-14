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
	"github.com/compose/gocomposeapi"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Delete: resourceDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of new deployment (must be unique per account)",
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Type:        schema.TypeString,
				Description: "The account in which the deployment will be placed",
				Required:    true,
				ForceNew:    true,
			},
			"datacenter": {
				Type:        schema.TypeString,
				Description: "Datacenter to use for Compose Hosted (this parameter is ignored if cluster_id is specified)",
				Optional:    true,
				ForceNew:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of deployment (mongodb, redis, postgresql, elastic_search, rethink, rabbitmq, etcd, scylla, mysql, disque)",
				Required:    true,
				ForceNew:    true,
			},
			"version": {
				Type:        schema.TypeString,
				Description: "Version of the software to deploy",
				Optional:    true,
				ForceNew:    true,
				Default:     "",
			},
			"units": {
				Type:        schema.TypeInt,
				Description: "Number of resource units to allocate to the deployment",
				Optional:    true,
				ForceNew:    true,
				Default:     "",
			},
		},
	}
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*composeapi.Client)
	deploymentParams := composeapi.DeploymentParams{
		Name:         d.Get("name").(string),
		AccountID:    d.Get("account_id").(string),
		Datacenter:   d.Get("datacenter").(string),
		DatabaseType: d.Get("type").(string),
		Version:      d.Get("version").(string),
		Units:        d.Get("units").(int),
	}

	deployment, errs := client.CreateDeployment(deploymentParams)
	if errs != nil {
		return concatErrors(errs)
	}

	d.SetId(deployment.ID)
	return resourceDeploymentRead(d, m)
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*composeapi.Client)
	deployment, errs := client.GetDeployment(d.Id())

	if errs != nil {
		return concatErrors(errs)
	}

	if err := d.Set("name", deployment.Name); err != nil {
		return err
	}
	if err := d.Set("type", deployment.Type); err != nil {
		return err
	}
	if err := d.Set("version", deployment.Version); err != nil {
		return err
	}

	return nil
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*composeapi.Client)

	_, errs := client.DeprovisionDeployment(d.Id())
	if errs != nil {
		return concatErrors(errs)
	}

	return nil
}

func concatErrors(errs []error) error {
	var result error
	for _, e := range errs {
		result = multierror.Append(result, e)
	}
	return result
}
