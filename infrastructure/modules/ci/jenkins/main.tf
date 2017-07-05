# VPC

# DevOps VPC

module "jenkins_vpc" {
  source = "../components/vpc"

  vpc_name           = "${var.environment}"
  cidr               = "${var.network_cidr}"
  private_subnets    = 0
  public_subnets     = 3
  enable_nat_gateway = false
}

# Jenkins Host


#

