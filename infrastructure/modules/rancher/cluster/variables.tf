variable "rancher_host_url" {
  description = "Rancher Server URL"
}

variable "rancher_host_token" {
  description = "Rancher agent environment token"
}

variable "vpc_id" {
  description = "VPC ID to add resources in"
}

variable "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
}

variable "subnet_ids" {
  type = "list"
  description = "The subnets to launch Rancher agents in"
}

variable "rancher_asg_min_size" {
  description = "The Rancher auto scaling group minimal size"
  default = 1
}

variable "rancher_asg_max_size" {
  description = "The Rancher auto scaling group minimal size"
  default = 4
}

variable "rancher_asg_desired_capacity" {
  description = "The Rancher auto scaling group minimal size"
  default = 1
}

variable "rancher_key_name" {
  description = "The Rancher auto scaling group minimal size"
}

variable "ami" {
  default = "ami-2d39803a"
}

variable "security_groups" {
  type = "list"
  default = []
}
