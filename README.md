# terraform-provider-compose

A [Terraform](https://www.terraform.io) plugin for managing [IBM Compose](https://compose.com/).

## Dependencies

This provider uses [gocomposeapi](https://github.com/compose/gocomposeapi) which is a wrapper around 
the [Compose REST API](https://apidocs.compose.com).

## Provider Configuration

```
provider "compose" {
  api_token = <COMPOSE_API_TOKEN>
}
```

| Property            | Description                                       
| ----------------    | -----------------------                              
| `api_token`         | The API token generated from the Compose UI console  

## Resources
### `deployment`

A resource for managing Compose deployments.

#### Example

```
resource "compose_deployment" "new-deployment" {
  name = "deployment_v1"
  account_id = <COMPOSE_ACCOUNT_ID>
  datacenter = "aws:us-west-2"
  type = "rabbitmq"
  version = "3.7.9"
}
```

#### Properties

| Property              | Description                                                        
| ----------------      | ----------------------                                           
| `name`                | The name of the deployment                                        
| `account_id`          | The ID of the account in which the deployment will be created    
| `datacenter`          | Datacenter to deploy to                                            
| `type`                | Type of deployment         
| `version`             | Version of software 
