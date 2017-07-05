
variable "module_source" {
  description "Module Source"
  default = "{{.Variables.moduleSource}}/infrastructure/modules/vpc"
}

variable "environment" {
  description "Name of Environment"
  default = "{{.Variables.environment}}"
}

variable "network_cidr" {
  description "CIDR for Network"
  default = "{{.Variables.networkCidr}}"
}

variable "aws_region" {
  description "AWS Region"
  default = "{{.Variables.awsRegion}}"
}

variable "aws_profile" {
  description "AWS IAM Profile"
  default = "{{.Variables.awsProfile}}"
}
