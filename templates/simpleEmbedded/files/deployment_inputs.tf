variable "region" {}

variable "profile" {}

variable "ssh_public_key_path" {
  default = "~/.ssh/id_rsa.pub"
}

variable "environment" {
  default = "tools"
}

variable "stack" {
  default = "server"
}

variable "network_cidr" {
  default = "10.0.0.0/16"
}
