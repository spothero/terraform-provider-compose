variable "compose_api_token" {}
variable "compose_account_id" {}

provider "compose" {
  api_token = "${var.compose_api_token}"
}

resource "compose_deployment" "my-deployment" {
  name = "deployment_v1"
  account_id = "${var.compose_account_id}"
  datacenter = "aws:us-west-2"
  type = "rabbitmq"
  version = "3.7.9"
  units = 2
}
