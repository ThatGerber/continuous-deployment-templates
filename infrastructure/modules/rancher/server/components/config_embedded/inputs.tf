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
