// Server Settings

variable "ami_id" {
  description = "The Ubuntu AMI to launch"
}

variable "key_name" {
  description = "The EC2 KeyPair to associate to the instances"
  default     = ""
}

variable "instance_profile" {
  description = "The EC2 IAM Instance Profile to associate to instanes"
  default     = ""
}

variable "instance_type" {
  description = "The EC2 Instance Type to launch"
  default     = "t2.micro"
}

variable "vpc_id" {}

variable "security_group_ids" {
  type    = "list"
  default = []
}

variable "has_public_ip" {
  default = "true"
}

variable "config" {}

variable "root_volume_size" {
  default = "32"
}

variable "environment" {
  description = "The environment (devPhase) to name this instance. Part of the Netflix Frigga naming pattern"
  default     = "tools"
}

variable "stack" {
  description = "The stack name for this instance. Part of the Netflix Frigga naming pattern"
  default     = "server"
}

variable "port" {
  description = "The port Rancher is exposed on. Needs to match that in the configuration"
  default     = "8080"
}

// Group Settings
variable "asg_max_size" {
  description = "The max # of servers to run"
  default     = "1"
}

variable "asg_min_size" {
  description = "The min # of servers to run"
  default     = "1"
}

variable "asg_desired_capacity" {
  description = "The desired # of servers to run"
  default     = "1"
}

variable "asg_subnet_ids" {
  description = "The subnets to launch instances in"
  type        = "list"
}

// ELB Settings
variable "elb_subnet_ids" {
  description = "The subnet to launch the ELB in"
  type        = "list"
}

variable "elb_lb_port" {
  description = "The public port to expose on the ELB"
  default     = "80"
}

variable "elb_access_cidr" {
  description = "The default allowable IPs to allow access (SSH, WEB)"
  default     = "0.0.0.0/0"
}

variable "ebl_security_group_ids" {
  description = "Addition Security Groups to add to the ELB"
  type        = "list"
  default     = []
}
