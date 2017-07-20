variable "environment" {
  description = "Name of Environment"
  default = "{{.Var.environment}}"
}

variable "network_cidr" {
  description = "CIDR for Network"
  default = "{{.Var.networkCidr}}"
}
