# Terraform State Backend
terraform {
  required_version = "< 0.9"
  {{if (eq .Var.tfBackend "s3")}}
  backend "s3" {}
  {{else}}
  backend "local" {}
  {{end}}
}

# AWS provider
provider "aws" {
  region  = "${var.aws_region}"

  profile = "${var.aws_profile}"
}
