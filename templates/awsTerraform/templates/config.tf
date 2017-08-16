variable "environment" {
  description = "Name of Environment"
  default     = "{{.Var.environment}}"
}

variable "network_cidr" {
  description = "CIDR for Network"
  default     = "{{.Var.networkCidr}}"
}

variable "aws_region" {
  description = "AWS Region"
  default     = "{{.Var.awsRegion}}"
}

variable "aws_profile" {
  description = "AWS IAM Profile"
  default     = "{{.Var.awsProfile}}"
}
