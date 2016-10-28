// Global Settings

variable "environment" {
  description = "The environment (devPhase) to name this instance. Part of the Netflix Frigga naming pattern"
  default = "tools"
}

variable "stack" {
  description = "The stack name for this instance. Part of the Netflix Frigga naming pattern"
  default = "server"
}

// Config Settings
variable "image" {
  description = "The docker image containing the Rancher server application"
  default = "rancher/server:latest"
}

variable "port" {
  description = "The port to expose Rancher on"
  default = "8080"
}

variable "docker_version" {
  description = "The version of Docker to install on the host"
  default = "1.11.2-0"
}

variable "ubuntu_version" {
  description = "The version of Ubunut to install"
  default = "xenial"
}

// Server Settings

variable "key_name" {
  description = "The EC2 KeyPair to associate to the instance"
  default = ""
}

variable "instance_profile" {
  description = "The EC2 IAM Instance Profile to associate to instane"
  default = ""
}

variable "instance_type" {
  description = "The EC2 Instance Type to launch"
  default = "t2.micro"
}

variable "access_cidr" {
  description = "The default allowable IPs to allow access (SSH, WEB)"
  default = "0.0.0.0/0"
}

variable "security_group_ids" {
  type = "list"
  default = []
}

variable "has_public_ip" {
  default = "true"
}

variable "root_volume_size" {
  default = "32"
}
