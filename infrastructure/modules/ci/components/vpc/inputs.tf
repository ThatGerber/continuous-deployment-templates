variable "vpc_name" {
  default = "devops"
}

variable "environment" {
  default = "devops"
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  default     = "10.0.0.0/16"
}

variable "num_public_subnets" {
  description = "Num. of Public Subnets"
  default     = 3
}

variable "num_private_subnets" {
  description = "Num. of Private Subnets"
  default     = 3
}

variable "use_nat_gateway" {
  default = true
}

variable "use_dns_hostnames" {
  default = true
}

variable "enable_dns_support" {
  default = true
}
