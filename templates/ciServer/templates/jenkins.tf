
# DevOps VPC

module "devops_vpc" {
  source = "{{.Variables.moduleSource}}/infrastructure/modules/vpc"

  environment = "{{.Variables.environment}}"
  vpc_prefix = "{{.Variables.networkCidr}}"
}
