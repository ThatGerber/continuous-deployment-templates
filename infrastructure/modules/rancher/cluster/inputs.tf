variable "environment" {}
variable "type" {
  description = "Rancher host type. Will be added as a host label of host.type={value}"
  default = "app"
}

// Rancher VPC Variables
variable "vpc_id" {
  description = "VPC ID to add resources in"
}

// Rancher ASG Variables
variable "asg_subnet_ids" {
  type = "list"
  description = "The subnets to launch Rancher agents in"
}

variable "asg_min_size" {
  description = "The Rancher auto scaling group minimal size"
  default = 1
}

variable "asg_max_size" {
  description = "The Rancher auto scaling group minimal size"
  default = 4
}

variable "asg_desired_capacity" {
  description = "The Rancher auto scaling group minimal size"
  default = 1
}

variable "asg_load_balancers" {
  type = "list"
  default = []
}

// Rancher Launch Configuration Variables
variable "host_instance_type" {
  default = "m4.large"
}

variable "host_ami_id" {
  default = "ami-2d39803a"
}

variable "host_key_name" {
  description = "The Rancher auto scaling group minimal size"
}

variable "host_iam_instance_profile" {
  default = ""
}

variable "host_additional_security_group_ids" {
  type = "list"
  default = []
}

variable "host_root_volume_size" {
  default = "32"
}

variable "host_root_volume_type" {
  default = "gp2"
}


// Rancher Variables
variable "rancher_agent_image" {
  default = "rancher/agent:latest"
}

variable "rancher_api_url" {}
variable "rancher_api_version" {
  default = "v1"
}
variable "rancher_environment_token" {}
