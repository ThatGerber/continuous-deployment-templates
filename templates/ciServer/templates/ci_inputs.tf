variable "environment" {
  description = "Name of Environment"
  default = "{{.Variables.environment}}"
}

variable "network_cidr" {
  description = "CIDR for Network"
  default = "{{.Variables.networkCidr}}"
}
