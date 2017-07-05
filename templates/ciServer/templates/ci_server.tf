# Terraform State Backend
terraform {
  backend {

  }
}

# AWS provider
provider "aws" {
  region = "${var.aws_region}"
  profile = "${var.aws_profile}"
}

# CI Server
module "ci_server" {
  source       = "${var.module_source}/infrastructure/modules/ci/{{.Variables.ciType}}"
  name         = "ci-server"
  environment  = "${var.environment}"
  network_cidr = "${var.network_cidr}"
}
