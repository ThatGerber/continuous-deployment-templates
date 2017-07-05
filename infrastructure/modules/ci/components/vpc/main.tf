# VPC

data "external" "public_subnets" {
  program = ["python", "${path.module}/tools/generate_cidrs.py"]

  query = {
    vpc     = "${var.vpc_cidr}"
    subnets = "${var.num_public_subnets}"
    vis     = "public"
  }
}

data "external" "private_subnets" {
  program = ["python", "${path.module}/tools/generate_cidrs.py"]

  query = {
    vpc     = "${var.vpc_cidr}"
    subnets = "${var.num_public_subnets}"
    vis     = "private"
  }
}

data "aws_availability_zones" "available" {}

module "devops_vpc" {
  source = "github.com/terraform-community-modules/tf_aws_vpc"

  name = "${var.vpc_name}"

  cidr            = "${var.vpc_cidr}"
  private_subnets = "${values(data.external.private_subnets.result)}"
  public_subnets  = "${values(data.external.public_subnets.result)}"

  enable_nat_gateway   = "${var.use_nat_gateway}"
  enable_dns_hostnames = "${var.use_dns_hostnames}"
  enable_dns_support   = "${var.enable_dns_support}"

  azs = ["${data.aws_availability_zones.available.names}"]

  tags {
    "Terraform"   = "true"
    "Environment" = "${var.environment}"
  }
}
