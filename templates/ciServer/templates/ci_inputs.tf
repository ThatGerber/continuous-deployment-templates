variable "environment" {
  description = "Name of Environment"
  default = "{{.environment}}"
}

variable "network_cidr" {
  description = "CIDR for Network"
  default = "{{.networkCidr}}"
}
