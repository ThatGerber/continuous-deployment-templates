// Global Settings
variable "environment" {
  description = "The environment (devPhase) to name this instance. Part of the Netflix Frigga naming pattern"
  default = "tools"
}

variable "stack" {
  description = "The stack name for this instance. Part of the Netflix Frigga naming pattern"
  default = "server"
}

// Config Setings
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

variable "db_schema" {
  description = "The database schema for Rancher tables"
  default = "rancher"
}

variable "db_username" {
  description = "The username to access the database"
  default = "rancher"
}

variable "db_password" {
  description = "The password to access the database"
}

variable "db_url" {
  description = "The URL for the database"
}

variable "db_port" {
  description = "The database connection port"
  default = "3306"
}
