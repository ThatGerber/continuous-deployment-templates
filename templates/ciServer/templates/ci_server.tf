# Terraform State Backend
terraform {
  required_version = "< 0.9"

{{if (eq .Variables.tfBackend "s3")}}
  backend "s3" {
    bucket = "{{.Variables.tfStateBucket}}"
    key    = "infrastructure/terraform.tfstate"
    region = "{{.Variables.tfStateRegion}}"
  }{{else}}
  backend "local" {
    # Will need to be configured for S3.
    path = ".terraform/terraform.tfstate"
  }{{end}}
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
