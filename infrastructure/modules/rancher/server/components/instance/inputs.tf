// Server Settings

variable "ami_id" {
  description = "The Ubuntu AMI to launch"
}

variable "key_name" {
  description = "The EC2 KeyPair to associate to the instance"
  default     = ""
}

variable "instance_profile" {
  description = "The EC2 IAM Instance Profile to associate to instane"
  default     = ""
}

variable "instance_type" {
  description = "The EC2 Instance Type to launch"
  default     = "t2.micro"
}

variable "access_cidr" {
  description = "The default allowable IPs to allow access (SSH, WEB)"
  default     = "0.0.0.0/0"
}

variable "vpc_id" {}

variable "security_group_ids" {
  type    = "list"
  default = []
}

variable "subnet_id" {}

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
